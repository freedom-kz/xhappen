package app

import (
	"context"
	"net"

	"github.com/go-kratos/kratos/v2/log"
)

type Boss struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	loggger   log.Logger

	tcpListener net.Listener

	hubs []*Hub

	exitChan  chan int
	waitGroup util.WaitGroupWrapper
}
