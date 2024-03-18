package client

import (
	"context"
	"fmt"
	"xhappen/app/xcache/internal/conf"
	"xhappen/pkg/utils"

	router_pb "xhappen/api/router/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
)

type XcacheClient struct {
	conns     map[string]*grpc.ClientConn //服务ID对应的grpc client
	localHost string
	log       *log.Helper
}

func NewXcacheClient(conf *conf.Bootstrap, logger log.Logger) (*XcacheClient, func(), error) {
	host, err := utils.GetLocalIp()
	if err != nil {
		return nil, func() {}, err
	}

	client := &XcacheClient{
		log:       log.NewHelper(log.With(logger, "module", "client/xcache")),
		localHost: host,
		conns:     make(map[string]*grpc.ClientConn, 0),
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
	//构建新的存储结构
	conns := make(map[string]*grpc.ClientConn)
	for _, service := range services {
		md := service.Metadata
		endpointHost, ok := md["host"]
		if !ok || endpointHost == xClient.localHost {
			xClient.log.Log(log.LevelError, "msg", "registry service instance info missing or local", "info", service)
			continue
		}

		if exist, ok := xClient.conns[endpointHost]; !ok {
			conn, err := transgrpc.DialInsecure(
				context.Background(),
				transgrpc.WithEndpoint(endpointHost),
				transgrpc.WithMiddleware(
					recovery.Recovery(),
				),
			)
			if err != nil {
				xClient.log.Log(log.LevelError, "msg", "grpc dial error", "error", err)
				continue
			}
			conns[endpointHost] = conn
		} else {
			conns[endpointHost] = exist
		}
	}

	//删除不需要的连接
	for host, conn := range xClient.conns {
		if _, ok := conns[host]; !ok {
			conn.Close()
		}
	}
	//进行存储替换
	xClient.conns = conns
	return true
}

func (xCacheClient *XcacheClient) UserDeviceBind(ctx context.Context, target string, req *router_pb.DeviceBindRequest) (*router_pb.DeviceBindReply, error) {
	conn, ok := xCacheClient.conns[target]
	if !ok {
		return nil, fmt.Errorf("target router server not alive")
	}

	peer := router_pb.NewRouterClient(conn)
	return peer.UserDeviceBind(ctx, req)
}
