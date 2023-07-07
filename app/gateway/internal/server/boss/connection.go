package boss

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	protocol "xhappen/api/protocol/v1"
)

type Connection struct {
	net.Conn
	Boss        *Boss
	ConnectTime time.Time
	ClientId    string
	UserId      string
	Hostname    string
	Os          protocol.DeviceType
	UserType    protocol.UserType
	RoleType    protocol.RoleType
	Version     string

	deliver *protocol.Deliver

	KeepAlive    time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	MsgTimeout   time.Duration
	SyncEvery    time.Duration
	SampleRate   int32

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
		Hostname:       host,
		KeepAlive:      conf.Main.MinKeepAlive.AsDuration(),
		WriteTimeout:   conf.Main.WriteTimeout.AsDuration(),
		ReadTimeout:    conf.Main.MinKeepAlive.AsDuration(),
		MsgTimeout:     conf.Queue.MsgTimeout.AsDuration(),
		SyncEvery:      conf.Queue.MsgTimeout.AsDuration(),
		ReadyStateChan: make(chan int),
		ExitCh:         make(chan bool),
	}

	return connection
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

func (connection *Connection) Flush() error {

	var zeroTime time.Time
	if connection.WriteTimeout > 0 {
		connection.SetWriteDeadline(time.Now().Add(connection.WriteTimeout))
	} else {
		connection.SetWriteDeadline(zeroTime)
	}

	switch s := connection.Conn.(type) {
	case *TcpConn:
		return s.Writer.Flush()
	case *WsConn:
		return s.Writer.Flush()
	}
	return nil
}
