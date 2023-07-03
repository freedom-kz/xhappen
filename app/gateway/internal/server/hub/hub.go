package protocol

import (
	"context"
	"net"

	"github.com/go-kratos/kratos/v2/log"
)

type Hub struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	loggger   log.Logger

	tcpListener net.Listener
}
