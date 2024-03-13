package client

import (
	"context"
	"fmt"
	"xhappen/app/xcache/internal/conf"

	router_pb "xhappen/api/router/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
)

type XcacheClient struct {
	conns map[string]*grpc.ClientConn //服务ID对应的grpc client
	log   *log.Helper
}

func NewXcacheClient(conf *conf.Bootstrap, logger log.Logger) (*XcacheClient, func(), error) {

	client := &XcacheClient{
		log: log.NewHelper(log.With(logger, "module", "client/xcache")),
	}

	cleanUp := func() {
		for _, conn := range client.conns {
			conn.Close()
		}
	}

	return client, cleanUp, nil
}

// 集群节点变化变更
func (xClient *XcacheClient) update(services []*registry.ServiceInstance) bool {
	for _, service := range services {
		md := service.Metadata
		endpointAddr, ok := md["endpointAddr"]
		if !ok {
			continue
		}
		if _, ok := xClient.conns[endpointAddr]; !ok {
			conn, err := transgrpc.DialInsecure(
				context.Background(),
				transgrpc.WithEndpoint(endpointAddr),
				transgrpc.WithMiddleware(
					recovery.Recovery(),
				),
			)
			if err != nil {
				continue
			}
			xClient.conns[endpointAddr] = conn
		}
	}
	return true
}

func (xCacheClient *XcacheClient) UserDeviceBind(ctx context.Context, serverAddr string, req *router_pb.DeviceBindRequest) (*router_pb.DeviceBindReply, error) {
	conn, ok := xCacheClient.conns[serverAddr]
	if !ok {
		return nil, fmt.Errorf("target router server not alive")
	}

	peer := router_pb.NewRouterClient(conn)
	return peer.UserDeviceBind(ctx, req)
}
