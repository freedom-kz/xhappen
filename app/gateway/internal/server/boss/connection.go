package boss

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	pb "xhappen/api/business/v1"
	protocol "xhappen/api/protocol/v1"
	"xhappen/app/gateway/internal/packets"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	DEFAULT_DAIL_TIMEOUT  = 3 * time.Second
	DEFAULT_RESEND_TICKER = 10 * time.Second
	DEFAULT_READ_TIMEOUT  = 10 * time.Second //默认读超时
	HandshakeTimeout      = 10 * time.Second //websocket握手超时
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

	deliver chan *protocol.Deliver
	sync    chan *protocol.Sync
	action  chan *protocol.Action

	KeepAlive    time.Duration
	DailTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	MsgTimeout   time.Duration
	FlushEvery   time.Duration
	SampleRate   int32

	writeLock      sync.RWMutex
	Closer         sync.Once
	ReadyStateChan chan int
	ExitCh         chan bool

	ReadyCount    int64
	InFlightCount int64
	FinishCount   uint64
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
		Conn:           conn,
		Boss:           boss,
		logger:         boss.loggger,
		Hostname:       host,
		ConnectTime:    time.Now(),
		KeepAlive:      conf.Main.MinKeepAlive.AsDuration(),
		DailTimeout:    DEFAULT_DAIL_TIMEOUT,
		WriteTimeout:   conf.Main.WriteTimeout.AsDuration(),
		ReadTimeout:    conf.Main.MinKeepAlive.AsDuration(),
		MsgTimeout:     conf.Queue.MsgTimeout.AsDuration(),
		FlushEvery:     conf.Queue.MsgTimeout.AsDuration(),
		ReadyStateChan: make(chan int),
		ExitCh:         make(chan bool),
	}

	return connection
}

func (connection *Connection) IOLoop() error {

	err := connection.processBind()
	if err != nil {
		return err
	}

	err = connection.processAuth()
	if err != nil {
		return err
	}

	messagePumpStartedChan := make(chan bool)
	go connection.messagePump(messagePumpStartedChan)
	<-messagePumpStartedChan
	connection.packetProcess()
	return nil
}

func (connection *Connection) processBind() error {
	if connection.DailTimeout > 0 {
		connection.SetReadDeadline(time.Now().Add(connection.DailTimeout))
	} else {
		connection.SetReadDeadline(time.Now().Add(DEFAULT_DAIL_TIMEOUT))
	}

	packet, err := connection.ReadPacket()

	if err != nil {
		connection.Boss.loggger.Log(log.LevelError, "msg", "read packet err", "err", err)
		return err
	}

	bind, ok := packet.(*packets.BindPacket)
	if !ok {
		connection.Boss.loggger.Log(log.LevelError, "msg", "read packet not bind", "body", packet.String())
		return err
	}

	bindAck := &packets.BindAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.BINDACK},
		BindAck: protocol.BindAck{
			BindRet:         true,
			ServerTimeStamp: uint64(time.Millisecond),
		},
	}

	in := &pb.BindRequest{
		ServerID: connection.Boss.serverId,
		BindInfo: &bind.Bind,
	}

	reply, err := connection.Boss.passClient.Bind(context.Background(), in)
	if err != nil {
		bindAck.BindRet = false
		bindAck.Err = &errors.FromError(err).Status
	} else {
		bindAck.BindRet = reply.Ret
		bindAck.Err = reply.Err
	}

	err = connection.Write(bindAck)
	if err != nil {
		connection.logger.Log(log.LevelDebug, "socket write fail", "clientId", connection.ClientId, "err", err)
		return err
	}
	if bindAck.BindRet == false {
		connection.logger.Log(log.LevelDebug, "msg", "Bind reply fail", "clientId", connection.ClientId, "err", reply.Err)
		connection.Flush()
		return fmt.Errorf(bindAck.Err.Reason)
	}
	connection.ClientId = bind.ClientID
	connection.KeepAlive = time.Duration(bind.KeepAlive)
	connection.Version = strconv.Itoa(int(bind.CurVersion))
	connection.Os = bind.DeviceType
	connection.ReadyCount = int64(bind.QueueSize)
	connection.LoginType = bind.LoginType
	return nil
}

func (connection *Connection) processAuth() error {
	if connection.ReadTimeout > 0 {
		connection.SetReadDeadline(time.Now().Add(connection.ReadTimeout))
	} else {
		connection.SetReadDeadline(time.Now().Add(DEFAULT_READ_TIMEOUT))
	}

	packet, err := connection.ReadPacket()

	if err != nil {
		connection.Boss.loggger.Log(log.LevelError, "msg", "read packet err", "err", err)
		return err
	}

	auth, ok := packet.(*packets.AUTHPacket)
	if !ok {
		connection.Boss.loggger.Log(log.LevelError, "msg", "read packet not bind", "body", packet.String())
		return err
	}

	ack := &packets.AuthAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.AUTHACK},
		AuthAck: protocol.AuthAck{
			AuthRet: true,
		},
	}

	in := &pb.AuthRequest{
		ClientId: connection.ClientId,
		AuthInfo: &auth.Auth,
	}

	reply, err := connection.Boss.passClient.Auth(context.Background(), in)
	if err != nil {
		ack.AuthRet = false
		ack.Err = &errors.FromError(err).Status
	} else {
		ack.AuthRet = reply.Ret
		ack.Uid = reply.Uid
		ack.Err = reply.Err
	}

	err = connection.Write(auth)
	if err != nil {
		connection.logger.Log(log.LevelDebug, "socket write fail", "clientId", connection.ClientId, "err", err)
		return err
	}
	if ack.AuthRet == false {
		connection.logger.Log(log.LevelDebug, "msg", "auth reply fail", "clientId", connection.ClientId, "err", reply.Err)
		connection.Flush()
		return fmt.Errorf(ack.Err.Reason)
	}

	connection.UserId = reply.Uid
	connection.RoleType = reply.Role
	connection.UserType = reply.UType
	connection.tokenExpire = reply.TokenExpire.AsTime()
	return nil
}

func (connection *Connection) messagePump(startedChan chan bool) {
	var err error
	var flusherChan <-chan time.Time
	var reSendChan <-chan time.Time

	dChan := connection.deliver
	aChan := connection.action
	sChan := connection.sync

	reSendTicker := time.NewTicker(DEFAULT_RESEND_TICKER)
	outputBufferTicker := time.NewTicker(connection.FlushEvery)
	reSendChan = reSendTicker.C
	flushed := true
	close(startedChan)
	for {
		if flushed {
			flusherChan = nil
		} else {
			flusherChan = outputBufferTicker.C
		}

		select {
		case <-flusherChan:
		case <-dChan:
		case <-aChan:
		case <-sChan:
		case <-reSendChan:
		case <-connection.ExitCh:
			goto exit
		}
	}
exit:
	reSendTicker.Stop()
	outputBufferTicker.Stop()
	connection.Shutdown(true)
	if err != nil {
		connection.logger.Log(log.LevelError, "msg", "send goroutine exit", "clientId", connection.ClientId, "userId", connection.UserId, "err", err)
	}
}

func (connection *Connection) packetProcess() error {
	var err error
	var packet packets.ControlPacket
	for {
		packet, err = connection.ReadPacket()
		if err != nil {
			break
		}

		err = connection.exec(packet)
		if err != nil {
			break
		}
	}

	close(connection.ExitCh)
	connection.Shutdown(true)
	return nil
}

func (connection *Connection) exec(packet packets.ControlPacket) error {
	var err error
	switch pkt := packet.(type) {
	case *packets.SubmitPacket:
	case *packets.SyncAckPacket:
	case *packets.DeliverAckPacket:
	case *packets.ActionPacket:
	case *packets.ActionAckPacket:
	case *packets.PingPacket:
	case *packets.QuitPacket:
	default:
		err = fmt.Errorf("invalid message type %s.", pkt.String())
	}
	return err
}

func (connection *Connection) ReadPacket() (cp packets.ControlPacket, err error) {
	return packets.ReadPacket(connection)
}

// 客户端远程网络地址
func (connection *Connection) String() string {
	return connection.RemoteAddr().String()
}

func (connection *Connection) IsReadyForMessages() bool {
	readyCount := atomic.LoadInt64(&connection.ReadyCount)
	inFlightCount := atomic.LoadInt64(&connection.InFlightCount)

	if inFlightCount >= readyCount || readyCount <= 0 {
		return false
	}

	return true
}

func (connection *Connection) tryUpdateReadyState() {
	// you can always *try* to write to ReadyStateChan because in the cases
	// where you cannot the message pump loop would have iterated anyway.
	// the atomic integer operations guarantee correctness of the value.
	select {
	case connection.ReadyStateChan <- 1:
	default:
	}
}

func (connection *Connection) FinishedMessage() {
	atomic.AddUint64(&connection.FinishCount, 1)
	atomic.AddInt64(&connection.InFlightCount, -1)
	connection.tryUpdateReadyState()
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

func (connection *Connection) Shutdown(push bool) {
	connection.Closer.Do(func() {
		connection.Conn.Close()
	})
}
