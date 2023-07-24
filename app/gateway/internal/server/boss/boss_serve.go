package boss

import (
	"net"

	"github.com/go-kratos/kratos/v2/log"
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
	connnection := newConnection(conn, bossServer.boss)
	err := connnection.IOLoop()

	if err != nil {
		bossServer.boss.logger.Log(log.LevelInfo, "msg", "socket io err.", "err", err, "hosname", connnection.String())
	}
}
