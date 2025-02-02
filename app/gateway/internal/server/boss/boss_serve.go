package boss

import (
	"net"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	ReadBufferSize  = 1024 * 8 //ws读缓冲大小
	WriteBufferSize = 1024 * 8 //ws写缓冲大小
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
		bossServer.boss.logger.Log(log.LevelInfo,
			"msg", "socket io err",
			"err", err,
			"deviceId", connnection.DeviceId,
			"SendBytes", connnection.SendBytes,
			"ReceiveBytes", connnection.ReceiveBytes,
			"hosname", connnection.String())
	} else {
		bossServer.boss.logger.Log(log.LevelInfo, "msg", "graceful quit", "hosname", connnection.String())
	}
}
