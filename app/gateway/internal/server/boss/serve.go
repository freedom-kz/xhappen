package boss

import (
	"net"
)

const (
	ReadBufferSize  = 1024 * 4 //ws读缓冲大小
	WriteBufferSize = 1024 * 4 //ws写缓冲大小
)

type ConnHandler interface {
	Handle(net.Conn)
}

type BossServer struct {
	boss *Boss
}

func (bossServer *BossServer) Handle(conn net.Conn) {

}

func (bossServer *BossServer) newConnection(conn net.Conn) *Connection {
	cli := &Connection{
		Conn:   conn,
		boss:   bossServer.boss,
		exitCh: make(chan bool),
	}
	return cli
}
