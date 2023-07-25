package boss

import (
	"fmt"
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
	DEFAULT_RESEND_TICKER = 10 * time.Second
	DEFAULT_READ_TIMEOUT  = 10 * time.Second //默认读超时
	HandshakeTimeout      = 10 * time.Second //websocket握手超时

	MESSAGE_RETRY_MAX = 3
)

const (
	STATE_BIND = iota
	STATE_AUTH
	STATE_SYNC
	STATE_NORMAL
	STATE_QUIT
)

type Connection struct {
	net.Conn
	logger log.Logger

	Boss        *Boss
	ConnectTime time.Time
	ClientId    string
	UserId      string
	tokenExpire time.Time
	Hostname    string
	Os          protocol.DeviceType
	UserType    protocol.UserType
	RoleType    protocol.RoleType
	LoginType   protocol.LoginType
	Version     string
	state       int

	deliverCh chan *protocol.Deliver
	syncCh    chan *protocol.Sync
	actionCh  chan *protocol.Action

	expectNextSequence uint64
	syncSessions       map[uint64]struct{}
	inFlightMutex      sync.Mutex
	inFlightMessages   map[uint64]*Message
	inFlightAPQ        inFlightAPqueue //发送队列
	inFlightAMutex     sync.Mutex
	inFlightAMessages  map[uint64]*AMessage
	inFlightPQ         inFlightPqueue //发送队列
	toFlightMutex      sync.Mutex
	toFlightMessages   map[uint64]*Message
	toFlightPQ         inFlightPqueue //缓冲队列

	KeepAlive    time.Duration
	DailTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	MsgTimeout   time.Duration
	FlushEvery   time.Duration
	SampleRate   int32

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

	connection := &Connection{
		Conn:             conn,
		Boss:             boss,
		logger:           boss.logger,
		state:            STATE_BIND,
		Hostname:         host,
		ConnectTime:      time.Now(),
		syncSessions:     make(map[uint64]struct{}),
		inFlightMessages: make(map[uint64]*Message),
		inFlightPQ:       newInFlightPqueue(100),
		toFlightMessages: make(map[uint64]*Message),
		toFlightPQ:       newInFlightPqueue(100),
		KeepAlive:        conf.Main.MinKeepAlive.AsDuration(),
		DailTimeout:      DEFAULT_DAIL_TIMEOUT,
		WriteTimeout:     conf.Main.WriteTimeout.AsDuration(),
		ReadTimeout:      conf.Main.MinKeepAlive.AsDuration(),
		MsgTimeout:       conf.Queue.MsgTimeout.AsDuration(),
		FlushEvery:       conf.Queue.MsgTimeout.AsDuration(),
		StateChan:        make(chan int),
		ReadyStateChan:   make(chan int),
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

func (connection *Connection) IsReadyForMessages() bool {
	if connection.state == STATE_QUIT {
		return false
	}

	readyCount := atomic.LoadInt64(&connection.ReadyCount)
	inFlightCount := atomic.LoadInt64(&connection.InFlightCount)

	if inFlightCount >= readyCount || readyCount <= 0 {
		return false
	}

	return true
}

func (connection *Connection) sendConnState(state int) {
	connection.StateChan <- state
}

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

func (connection *Connection) SetSampleRate(sampleRate int32) error {
	if sampleRate < 0 || sampleRate > 99 {
		return fmt.Errorf("sample rate (%d) is invalid", sampleRate)
	}
	atomic.StoreInt32(&connection.SampleRate, sampleRate)
	return nil
}

func (connection *Connection) processExpectSequence(sequence uint64) {
	if sequence >= connection.expectNextSequence {
		connection.expectNextSequence = sequence
	}
	connection.expectNextSequence++

}

func (connection *Connection) Write(packet packets.ControlPacket) error {
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

func (connection *Connection) Flush() error {
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

//业务关闭
func (connection *Connection) Shutdown(active bool) {
	connection.Closer.Do(func() {
		err := connection.Flush()
		if err != nil {
			connection.logger.Log(log.LevelDebug, "msg", "socket flush err.", "err", err, "hosname", connection.String())
		}
		if !active && connection.RoleType != protocol.RoleType_ROLE_CUSTOMER_SERVICE {
			//TODO,离线消息推送
		}

		connection.sendConnState(STATE_QUIT)

		err = connection.Conn.Close()
		if err != nil {
			connection.logger.Log(log.LevelInfo, "msg", "socket closed err.", "err", err, "hosname", connection.String())
		}
	})
}
