package client

import (
	"xhappen/app/xcache/internal/conf"
	// sequence_pb "xhappen/api/sequence"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"google.golang.org/grpc"
)

type XcacheClient struct {
	conns map[int]*grpc.ClientConn //服务ID对应的grpc client
	log   *log.Helper
}

func NewXcacheClient(conf *conf.Bootstrap, logger log.Logger) (*XcacheClient, func(), error) {
	return &XcacheClient{
		log: log.NewHelper(log.With(logger, "module", "client/xcache")),
	}, func() {}, nil
}

func (xClient *XcacheClient) update(services []*registry.ServiceInstance) bool {
	return true
}
