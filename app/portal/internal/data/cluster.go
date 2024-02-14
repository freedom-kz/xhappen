package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

const gatewayEndpoint = ""

type GatewayServiceRepo struct {
	log *log.Helper

	watch registry.Watcher
	r     registry.Discovery
	srvs  registry.ServiceInstance
}

func NewGatewayServiceRepo(ctx context.Context, logger log.Logger, r registry.Discovery) *GatewayServiceRepo {

	r.GetService(ctx, gatewayEndpoint)
	watch, err := r.Watch(ctx, gatewayEndpoint)
	if err != nil {
		logger.Log(log.LevelFatal, "msg", "watch gateway fatal", "err", err)
	}
	go func() {
		for {
			// srvs, err := watch.Next()
		}
	}()

	return &GatewayServiceRepo{
		r:     r,
		watch: watch,
		log:   log.NewHelper(logger),
	}
}
