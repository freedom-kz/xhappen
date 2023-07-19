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
	inFlightPQ         inFlightPqueue
	toFlightMessages   map[uint64]*Message
	toFlightPQ         inFlightPqueue

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
		ReadyStateChan:   make(chan int),
		ExitCh:           make(chan bool),
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
	err = connection.packetProcess()
	return err
}

func (connection *Connection) processBind() error {
	connection.state = STATE_BIND
	if connection.DailTimeout > 0 {
		connection.SetReadDeadline(time.Now().Add(connection.DailTimeout))
	} else {
		connection.SetReadDeadline(time.Now().Add(DEFAULT_DAIL_TIMEOUT))
	}

	packet, err := connection.ReadPacket()

	if err != nil {
		connection.logger.Log(log.LevelError, "msg", "read packet err", "err", err)
		return err
	}

	bind, ok := packet.(*packets.BindPacket)
	if !ok {
		connection.logger.Log(log.LevelError, "msg", "read packet not bind", "body", packet.String())
		return err
	}

	bindAck := &packets.BindAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.BINDACK},
		BindAck: protocol.BindAck{
			BindRet:         true,
			ServerTimeStamp: uint64(time.Now().UnixMilli()),
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
	if !bindAck.BindRet {
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
	connection.state = STATE_AUTH
	connection.state = STATE_AUTH
	if connection.ReadTimeout > 0 {
		connection.SetReadDeadline(time.Now().Add(connection.ReadTimeout))
	} else {
		connection.SetReadDeadline(time.Now().Add(DEFAULT_READ_TIMEOUT))
	}

	packet, err := connection.ReadPacket()

	if err != nil {
		connection.Boss.logger.Log(log.LevelError, "msg", "read packet err", "err", err)
		return err
	}

	auth, ok := packet.(*packets.AUTHPacket)
	if !ok {
		connection.Boss.logger.Log(log.LevelError, "msg", "read packet not bind", "body", packet.String())
		return err
	}

	ack := &packets.AuthAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.AUTHACK},
		AuthAck: protocol.AuthAck{
			AuthRet: true,
		},
	}

	in := &pb.AuthRequest{
		ClientId:    connection.ClientId,
		BindVersion: uint64(connection.ConnectTime.UnixNano()),
		AuthInfo:    &auth.Auth,
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
	if !ack.AuthRet {
		connection.logger.Log(log.LevelDebug, "msg", "auth reply fail", "clientId", connection.ClientId, "err", reply.Err)
		connection.Flush()
		return fmt.Errorf(ack.Err.Reason)
	}

	connection.UserId = reply.Uid
	connection.RoleType = reply.Role
	connection.UserType = reply.UType
	connection.tokenExpire = reply.TokenExpire.AsTime()
	for _, sid := range reply.Sessions {
		connection.syncSessions[sid] = struct{}{}
	}
	if len(connection.syncSessions) == 0 {
		connection.state = STATE_NORMAL
	} else {
		connection.state = STATE_SYNC
	}

	connection.Boss.AddConnToHub(connection)
	return nil
}

func (connection *Connection) processSubmit(submit *protocol.Submit) error {
	in := &pb.SubmitRequest{
		Userid:   connection.UserId,
		Clientid: connection.ClientId,
		Submit:   submit,
	}
	reply, err := connection.Boss.passClient.Submit(context.Background(), in)

	ack := &packets.SubmitAckPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.SUBMITACK},
		SubmitAck: protocol.SubmitAck{
			SubmitRet: true,
			Id:        submit.Id,
			Timestamp: uint64(time.Now().UnixMilli()),
		},
	}
	if err != nil {
		ack.SubmitRet = false
		ack.Err = &errors.FromError(err).Status
	} else {
		ack.SubmitRet = reply.Ret
		ack.Sequence = reply.Sequence
		ack.SessionId = reply.Sequence
		ack.Err = reply.Err
	}

	return nil
}

func (connection *Connection) processSyncAck(syncAck *protocol.SyncAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process SyncAck", "id", syncAck.Id)
	delete(connection.syncSessions, uint64(syncAck.Id))
	if len(connection.syncSessions) == 0 {
		packet := &packets.SyncConfirmPacket{FixedHeader: packets.FixedHeader{MessageType: packets.SYNCCONFIRM}}
		connection.Write(packet)
		connection.Flush()
	}
	//TODO,
	return nil
}

func (connection *Connection) processDeliverAck(deliverAck *protocol.DeliverAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process deliver ack")
	err := connection.FinishDeliver(deliverAck.Sequence)
	connection.logger.Log(log.LevelError, "msg", "process deliver ack", "error", err)
	return nil
}

func (connection *Connection) processAction(action *protocol.Action) error {
	return nil
}

func (connection *Connection) processActionAck(actionAck *protocol.ActionAck) error {
	connection.logger.Log(log.LevelDebug, "msg", "process action ack")
	err := connection.FinishAction(uint64(actionAck.Id))
	connection.logger.Log(log.LevelError, "msg", "process action ack", "error", err)
	return nil
}

func (connection *Connection) processPing(ping *protocol.Ping) error {
	pong := &packets.PongPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.SUBMITACK},
		Pong:        protocol.Pong{},
	}
	err := connection.Write(pong)
	if err != nil {
		return err
	}
	err = connection.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (connection *Connection) processQuit(quit *protocol.Quit) error {
	return fmt.Errorf("not error normal exit")
}

func (connection *Connection) messagePump(startedChan chan bool) {
	var err error
	var flusherChan <-chan time.Time
	var reSendChan <-chan time.Time

	var dChan chan *protocol.Deliver = nil

	aChan := connection.actionCh
	sChan := connection.syncCh

	reSendTicker := time.NewTicker(DEFAULT_RESEND_TICKER)
	outputBufferTicker := time.NewTicker(connection.FlushEvery)
	reSendChan = reSendTicker.C
	flushed := true
	close(startedChan)
	for {
		if !connection.IsReadyForMessages() {
			dChan = nil
		} else if flushed {
			flusherChan = nil
			dChan = connection.deliverCh
		} else {
			flusherChan = outputBufferTicker.C
			dChan = connection.deliverCh
		}

		select {
		case <-flusherChan:
			err = connection.Flush()
			connection.logger.Log(log.LevelError, "msg", "socket flush err")
		case d := <-dChan:
			msg := NewMessage(d)
			err := connection.StartInflight(msg)
			if err != nil {
				connection.logger.Log(log.LevelError, "msg", "in startInflight", "err", err)
			}
			err = connection.WriteDeliver(msg)
			if err != nil {
				goto exit
			}
			flushed = false
		case <-aChan:
		case <-sChan:

		case <-reSendChan:
		case <-connection.ReadyStateChan:
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
		select {
		case <-connection.ExitCh:
			goto exit
		default:
			packet, err = connection.ReadPacket()
			if err != nil {
				break
			}

			err = connection.exec(packet)
			if err != nil {
				goto exit
			}
		}
	}
exit:
	connection.Shutdown(true)
	return err
}

func (connection *Connection) exec(packet packets.ControlPacket) error {
	var err error
	switch pkt := packet.(type) {
	case *packets.SubmitPacket:
		err = connection.processSubmit(&pkt.Submit)
	case *packets.SyncAckPacket:
		err = connection.processSyncAck(&pkt.SyncAck)
	case *packets.DeliverAckPacket:
		err = connection.processDeliverAck(&pkt.DeliverAck)
	case *packets.ActionPacket:
		err = connection.processAction(&pkt.Action)
	case *packets.ActionAckPacket:
		err = connection.processActionAck(&pkt.ActionAck)
	case *packets.PingPacket:
		err = connection.processPing(&pkt.Ping)
	case *packets.QuitPacket:
		err = connection.processQuit(&pkt.Quit)
	default:
		err = fmt.Errorf("invalid message type %s", pkt.String())
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

	select {
	case <-connection.ExitCh:
		return false
	default:
	}

	readyCount := atomic.LoadInt64(&connection.ReadyCount)
	inFlightCount := atomic.LoadInt64(&connection.InFlightCount)

	if inFlightCount >= readyCount || readyCount <= 0 {
		return false
	}

	return true
}

func (connection *Connection) StartInflight(msg *Message) error {
	msg.pri = time.Now().UnixNano()
	err := connection.pushInFlightMessage(msg)
	if err != nil {
		return err
	}
	connection.addToInFlightPQ(msg)
	return nil
}

func (connection *Connection) addToInFlightPQ(msg *Message) {
	connection.inFlightMutex.Lock()
	connection.inFlightPQ.Push(msg)
	connection.inFlightMutex.Unlock()
}

func (connection *Connection) popInFlightMessage(sequence uint64) (*Message, error) {
	connection.inFlightMutex.Lock()
	msg, ok := connection.inFlightMessages[sequence]
	if !ok {
		connection.inFlightMutex.Unlock()
		return nil, fmt.Errorf("sequence not in flight")
	}
	delete(connection.inFlightMessages, sequence)
	connection.inFlightMutex.Unlock()
	return msg, nil
}

func (connection *Connection) removeFromInFlight(msg *Message) {
	connection.inFlightMutex.Lock()
	if msg.index == -1 {
		// this item has already been popped off the pqueue
		connection.inFlightMutex.Unlock()
		return
	}
	connection.inFlightPQ.Remove(msg.index)
	connection.inFlightMutex.Unlock()
}

func (connection *Connection) pushInFlightMessage(msg *Message) error {
	connection.inFlightMutex.Lock()
	_, ok := connection.inFlightMessages[msg.Sequence]
	if ok {
		connection.inFlightMutex.Unlock()
		return fmt.Errorf("sequence already in flight")
	}
	connection.inFlightMessages[msg.Sequence] = msg
	connection.inFlightMutex.Unlock()
	return nil
}

func (connection *Connection) WriteDeliver(message *Message) error {
	message.Attempts++
	message.pri = message.deliveryTS.Add(connection.MsgTimeout).UnixNano()
	if message.Attempts > MESSAGE_RETRY_MAX {
		return fmt.Errorf("messgae send > %d", MESSAGE_RETRY_MAX)
	}

	if message.Attempts > 1 {
		connection.logger.Log(log.LevelInfo, "msg", "message retry")
	}

	packet := &packets.DeliverPacket{
		FixedHeader: packets.FixedHeader{MessageType: packets.DELIVER},
		Deliver:     *message.Deliver,
	}
	return connection.Write(packet)
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

func (connection *Connection) SendDeliverCh(deliver *protocol.Deliver) {
	connection.deliverCh <- deliver
}

func (connection *Connection) SendSyncCh(sync *protocol.Sync) {
	connection.syncCh <- sync
}

func (connection *Connection) SendActionCh(action *protocol.Action) {
	connection.actionCh <- action
}

func (connection *Connection) FinishDeliver(sequence uint64) error {
	msg, err := connection.popInFlightMessage(sequence)
	if err != nil {
		return err
	}
	connection.removeFromInFlight(msg)
	connection.FinishedMessage()
	return nil
}
func (connection *Connection) FinishAction(sequence uint64) error {
	//TODO
	return nil
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

//业务关闭
func (connection *Connection) Shutdown(initiative bool) {
	connection.Closer.Do(func() {
		connection.Flush()
		if !initiative && connection.RoleType != protocol.RoleType_ROLE_CUSTOMER_SERVICE {
			//TODO,离线消息推送
		}

		connection.state = STATE_QUIT
		close(connection.ExitCh)
		err := connection.Conn.Close()

		if err != nil {
			connection.logger.Log(log.LevelInfo, "msg", "socket closed err.", "err", err, "hosname", connection.String())
		}
	})
}
