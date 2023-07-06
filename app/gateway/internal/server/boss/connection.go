package boss

import (
	"net"
	"sync"
	"time"

	protocol "xhappen/api/protocol/v1"
)

type Connection struct {
	net.Conn
	boss        *Boss
	ConnectTime time.Time
	clientId    string
	userId      string
	Os          protocol.DeviceType
	UserType    protocol.UserType
	roleType    protocol.RoleType
	version     string

	keepAlive    time.Duration
	writeTimeout time.Duration
	readTimeout  time.Duration
	msgTimeout   time.Duration
	syncEvery    time.Duration
	closer       sync.Once
	exitCh       chan bool
}
