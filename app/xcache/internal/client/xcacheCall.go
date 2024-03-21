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
	conns      map[int]*grpc.ClientConn //服务索引对应的grpc client
	localIndex int
	log        *log.Helper
}

func NewXcacheClient(conf *conf.Bootstrap, logger log.Logger) (*XcacheClient, func(), error) {
	client := &XcacheClient{
		log:        log.NewHelper(log.With(logger, "module", "client/xcache")),
		localIndex: -1,
		conns:      make(map[int]*grpc.ClientConn, 0),
	}

	cleanUp := func() {
		for _, conn := range client.conns {
			conn.Close()
		}
	}

	return client, cleanUp, nil
}

// 集群节点变化变更
func (xClient *XcacheClient) UpdateService(services map[int]*registry.ServiceInstance) bool {
	//构建新的存储结构
	conns := make(map[int]*grpc.ClientConn)
	for index, service := range services {
		//本机索引，不创建连接
		if index == xClient.localIndex {
			continue
		}
		//数据缺失
		endpointHost, ok := service.Metadata["host"]
		if !ok {
			continue
		}

		//新建/复用
		if exist, ok := xClient.conns[index]; !ok {
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
			conns[index] = conn
		} else {
			conns[index] = exist
		}
	}

	//删除不需要的连接
	for index, conn := range xClient.conns {
		if curConn, ok := conns[index]; !ok {
			conn.Close()
		} else if conn != curConn {
			conn.Close()
		}
	}
	//进行存储替换
	xClient.conns = conns
	return true
}

func (xCacheClient *XcacheClient) LocalStateModify(state bool, index int) {
	if state {
		xCacheClient.localIndex = index
	} else {
		xCacheClient.localIndex = -1
	}
}

func (xCacheClient *XcacheClient) UserDeviceBind(ctx context.Context, target int, req *router_pb.UserDeviceBindRequest) (*router_pb.UserDeviceBindReply, error) {
	conn, ok := xCacheClient.conns[target]
	if !ok {
		return nil, fmt.Errorf("target router server not alive")
	}
	peer := router_pb.NewRouterClient(conn)
	return peer.UserDeviceBind(ctx, req)
}
