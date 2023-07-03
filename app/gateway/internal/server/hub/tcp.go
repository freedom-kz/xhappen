package protocol

import (
	"net"

	"github.com/go-kratos/kratos/v2/log"
)

type tcpServer struct {
	hub *Hub
}

func (p *tcpServer) Handle(conn net.Conn) {
	p.hub.loggger.Log(log.LevelInfo, "msg", "TCP: new client", "remote", conn.RemoteAddr())

}
