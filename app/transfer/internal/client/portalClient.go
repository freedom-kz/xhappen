package client

import (
	"context"
	"os"

	pb "xhappen/api/portal/v1"
	"xhappen/app/gateway/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	srcgrpc "google.golang.org/grpc"
)

const endpoint = "discovery:///business"

type PortalClient struct {
	conn *srcgrpc.ClientConn
}

// 采用随机调用
func NewPortalClient(conf *conf.Bootstrap, logger log.Logger) (*PortalClient, func(), error) {
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
		transgrpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(r),
		transgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		logger.Log(log.LevelFatal, "msg", "transgrpc.DialInsecure fail", "err", err)
		os.Exit(1)
	}
	pb.NewUserClient(conn)

	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing the grpc client resources")
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}

	passClient := &PortalClient{
		conn: conn,
	}
	return passClient, cleanup, nil
}

// 设备绑定
func (portalClient *PortalClient) Bind(ctx context.Context, in *pb.BindRequest) (*pb.BindReply, error) {
	client := pb.NewUserHTTPClient(portalClient.conn)
	return client.Bind(ctx, in)
}
