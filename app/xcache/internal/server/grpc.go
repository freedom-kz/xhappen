package server

import (
	routerv1 "xhappen/api/router/v1"
	sequencev1 "xhappen/api/sequence/v1"
	"xhappen/app/xcache/internal/conf"
	"xhappen/app/xcache/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *conf.Bootstrap, sequence *service.SequenceService, router *service.RouterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}

	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	routerv1.RegisterRouterServer(srv, router)
	sequencev1.RegisterSequenceServer(srv, sequence)
	return srv
}
