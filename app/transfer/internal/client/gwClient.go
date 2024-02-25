package client

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	srcgrpc "google.golang.org/grpc"
)

/*
网关客户端管理
以IP为key存放连接
*/
const (
	SERVICE_NAME_GATEWAY = "gateway"
)

type GatewayClient struct {
	conns map[string]*srcgrpc.ClientConn
}

func NewGatewayClient(discovery registry.Discovery, logger log.Logger) (*GatewayClient, func(), error) {
	//同监听gateway，然后维护连接
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	services, err := discovery.GetService(ctx, SERVICE_NAME_GATEWAY)
	if err != nil {
		log.Fatal("msg", "failed connection to cluster: %v", err)
	}
	watcher, err := discovery.Watch(ctx, SERVICE_NAME_GATEWAY)
	if err != nil {
		log.Fatal("msg", "failed watch service[%s] in cluster: %v", SERVICE_NAME_GATEWAY, err)
	}

	gwClient := &GatewayClient{
		conns: make(map[string]*srcgrpc.ClientConn),
	}

}

// 更新服务客户端
func (repo *GatewayClient) update(services []*registry.ServiceInstance) {
	for _, service := range services {
		addr = ""

	}
}
