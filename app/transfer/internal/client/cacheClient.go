package client

import (
	"context"
	"os"

	"xhappen/app/transfer/internal/conf"

	pb "xhappen/api/router/v1"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	srcgrpc "google.golang.org/grpc"
)

const endpointXcache = "discovery:///xcache"

type XcahceClient struct {
	conn *srcgrpc.ClientConn
}

// 采用随机调用
func NewXcache(conf *conf.Bootstrap, logger log.Logger) (*XcahceClient, func(), error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Data.Etcd.Addr},
	})
	if err != nil {
		logger.Log(log.LevelFatal, "msg", "init etcd client fail", "err", err)
		os.Exit(1)
	}

	r := etcd.New(cli)

	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint(endpointXcache),
		grpc.WithDiscovery(r),
		transgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		logger.Log(log.LevelFatal, "msg", "transgrpc.DialInsecure fail", "err", err)
		os.Exit(1)
	}

	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing the grpc client resources")
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}

	xcacheClient := &XcahceClient{
		conn: conn,
	}
	return xcacheClient, cleanup, nil
}

// 设备绑定
func (xacheClient *XcahceClient) UserDeviceBind(ctx context.Context, in *pb.UserDeviceBindRequest) (*pb.UserDeviceBindReply, error) {
	client := pb.NewRouterClient(xacheClient.conn)

	return client.UserDeviceBind(ctx, in)
}

// 设备解绑
func (xacheClient *XcahceClient) UserDeviceUnBind(ctx context.Context, in *pb.UserDeviceUnBindRequest) (*pb.UserDeviceUnBindReply, error) {
	client := pb.NewRouterClient(xacheClient.conn)
	return client.UserDeviceUnBind(ctx, in)
}
