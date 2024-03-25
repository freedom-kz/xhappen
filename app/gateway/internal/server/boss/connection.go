package boss

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	protocol "xhappen/api/protocol/v1"
	"xhappen/app/gateway/internal/packets"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	DEFAULT_DAIL_TIMEOUT  = 3 * time.Second
	DEFAULT_RESEND_TICKER = 5 * time.Second
	DEFAULT_READ_TIMEOUT  = 10 * time.Second //默认读超时
	HandshakeTimeout      = 3 * time.Second  //websocket握手超时

	MESSAGE_ATTEMPTS_MAX = 3
)

const (
	STATE_INIT   = iota //初始化
	STATE_BIND          //设备绑定
	STATE_AUTH          //设备验证
	STATE_SYNC          //消息同步
	STATE_NORMAL        //普通收发
	STATE_QUIT          //退出
)

type Connection struct {
	net.Conn
	logger log.Logger

	Boss            *Boss
	ConnectTime     time.Time
	connectSequence uint64
	DeviceId        string
	UserId          string
	tokenExpire     time.Time
	Hostname        string
	Os              protocol.DeviceType
	UserType        protocol.UserType
	RoleType        protocol.RoleType
	LoginType       protocol.LoginType
	Version         int
	state           int

	deliverCh chan *protocol.Deliver
	syncCh    chan *protocol.Sync
	actionCh  chan *protocol.Action

	expectNextSequence uint64
	syncSessions       map[uint64]struct{}

	inFlightMutex     sync.Mutex
	inFlightMessages  map[uint64]*Message
	inFlightPQ        inFlightPqueue //deliver发送队列
	inFlightAMutex    sync.Mutex
	inFlightAPQ       inFlightAPqueue //action发送队列
	inFlightAMessages map[uint64]*AMessage
	toFlightMutex     sync.Mutex
	toFlightMessages  map[uint64]*Message
	toFlightPQ        inFlightPqueue //deliver缓冲队列

	KeepAlive    time.Duration
	DailTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	MsgTimeout   time.Duration
	FlushEvery   time.Duration

	writeLock      sync.RWMutex
	Closer         sync.Once
	StateChan      chan int
	ReadyStateChan chan int

	ReadyCount    int64
	InFlightCount int64  //包含action和deliver
	FinishCount   uint64 //包含action和deliver
	SendBytes     int64
	ReceiveBytes  int64
}

func newConnection(conn net.Conn, boss *Boss) *Connection {

	conf := boss.GetConfig()

	var host string
	if conn != nil {
		host, _, _ = net.SplitHostPort(conn.RemoteAddr().String())
	}

	connectSequence := atomic.AddUint64(&boss.GlobalConnectionSequence, 1)

	connection := &Connection{
		Conn:              conn,
		Boss:              boss,
		logger:            boss.logger,
		state:             STATE_INIT,
		Hostname:          host,
		connectSequence:   connectSequence,
		ConnectTime:       time.Now(),
		deliverCh:         make(chan *protocol.Deliver, 100),
		actionCh:          make(chan *protocol.Action, 100),
		syncCh:            make(chan *protocol.Sync, 100),
		syncSessions:      make(map[uint64]struct{}),
		inFlightMessages:  make(map[uint64]*Message),
		inFlightPQ:        newInFlightPqueue(100),
		toFlightMessages:  make(map[uint64]*Message),
		toFlightPQ:        newInFlightPqueue(100),
		inFlightAMessages: make(map[uint64]*AMessage),
		inFlightAPQ:       newInFlightAPqueue(100),
		KeepAlive:         conf.Main.MinKeepAlive.AsDuration(),
		DailTimeout:       DEFAULT_DAIL_TIMEOUT,
		WriteTimeout:      conf.Main.WriteTimeout.AsDuration(),
		ReadTimeout:       conf.Main.MinKeepAlive.AsDuration(),
		MsgTimeout:        conf.Queue.MsgTimeout.AsDuration(),
		FlushEvery:        conf.Queue.MsgTimeout.AsDuration(),
		StateChan:         make(chan int),
		ReadyStateChan:    make(chan int),
	}

	return connection
}

func (connection *Connection) IOLoop() error {
	messagePumpStartedChan := make(chan bool)
	go connection.messagePump(messagePumpStartedChan)
	<-messagePumpStartedChan

	err := connection.processBind()
	if err != nil {
		return err
	}
	err = connection.processAuth()
	if err != nil {
		return err
	}

	err = connection.packetProcess()
	return err
}

// 客户端远程网络地址
func (connection *Connection) String() string {
	return connection.RemoteAddr().String()
}

// 客户端接收状态判断
func (connection *Connection) IsReadyForMessages() bool {
	//状态匹配
	if connection.state != STATE_NORMAL {
		return false
	}

	readyCount := atomic.LoadInt64(&connection.ReadyCount)
	inFlightCount := atomic.LoadInt64(&connection.InFlightCount)

	if inFlightCount >= readyCount || readyCount <= 0 {
		return false
	}

	return true
}

// 发送客户端状态变更通知
func (connection *Connection) sendConnState(state int) {
	connection.StateChan <- state
}

// 触发客户端接收状态事件
func (connection *Connection) tryUpdateReadyState() {
	select {
	case connection.ReadyStateChan <- 1:
	default:
	}
}

func (connection *Connection) SendDeliverCh(deliver *protocol.Deliver) {
	connection.deliverCh <- deliver
}

func (connection *Connection) SendSyncCh(sync *protocol.Sync) {
	connection.syncCh <- sync
}

func (connection *Connection) SendActionCh(action *protocol.Action) {
	connection.actionCh <- action
}

func (connection *Connection) SendingMessage() {
	atomic.AddInt64(&connection.InFlightCount, 1)
}

// 对期望序列号进行加工
func (connection *Connection) processExpectSequence(sequence uint64) {
	if sequence >= connection.expectNextSequence {
		connection.expectNextSequence = sequence
	}
	connection.expectNextSequence++

}

// socket数据写入
func (connection *Connection) Write(packet packets.ControlPacket) error {
	//TODO，目前还是有锁，待优化
	connection.writeLock.Lock()
	defer connection.writeLock.Unlock()
	var zeroTime time.Time
	if connection.WriteTimeout > 0 {
		connection.SetWriteDeadline(time.Now().Add(connection.WriteTimeout))
	} else {
		connection.SetWriteDeadline(zeroTime)
	}
	err := packet.Write(connection.Conn)
	if err != nil {
		return err
	}
	return nil
}

// 数据缓冲刷新
func (connection *Connection) Flush() error {
	//目前还是有锁，待优化
	connection.writeLock.Lock()
	defer connection.writeLock.Unlock()
	var zeroTime time.Time
	if connection.WriteTimeout > 0 {
		connection.SetWriteDeadline(time.Now().Add(connection.WriteTimeout))
	} else {
		connection.SetWriteDeadline(zeroTime)
	}

	switch s := connection.Conn.(type) {
	case *TcpConn:
		return s.Flush()
	case *WsConn:
		return s.Flush()
	}
	return nil
}

// 业务关闭
func (connection *Connection) Shutdown(active bool) {
	connection.Closer.Do(func() {
		err := connection.Flush()
		if err != nil {
			connection.logger.Log(log.LevelDebug, "msg", "socket flush err.", "err", err, "hosname", connection.String())
		}

		if active {
			//预期内关闭（正常收发，多为客户端主动或网络情况触发）
			connection.sendConnState(STATE_QUIT)
		} else {
			//预期外关闭（踢下线，服务器关闭等业务执行，少数情况）
			err = connection.Conn.Close()
			if err != nil {
				connection.logger.Log(log.LevelInfo, "msg", "socket closed err.", "err", err, "hosname", connection.String())
			}
		}
		//预期内关闭的，看后面看是否需要补一次离线消息业务
		if active && connection.RoleType != protocol.RoleType_ROLE_CUSTOMER_SERVICE {
			//TODO,离线消息推送
		}
	})
}
