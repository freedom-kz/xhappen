package client

import (
	"context"
	"os"

	pb "xhappen/api/portal/v1"
	"xhappen/app/transfer/internal/conf"

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

// 设备Auth
func (portalClient *PortalClient) TokenAuth(ctx context.Context, in *pb.TokenAuthRequest) (*pb.TokenAuthReply, error) {
	client := pb.NewUserClient(portalClient.conn)
	return client.TokenAuth(ctx, in)
}

// 配置获取
func (portalClient *PortalClient) GetSocketHostConfig(ctx context.Context, in *pb.GetSocketHostConfigRequest) (*pb.GetSocketHostConfigReply, error) {
	client := pb.NewConfigClient(portalClient.conn)
	return client.GetSocketHostConfig(ctx, in)
}

// 获取用户信息
func (portalClient *PortalClient) GetSelfProfile(ctx context.Context, in *pb.GetSelfProfileRequest) (*pb.GetSelfProfileReply, error) {
	client := pb.NewUserClient(portalClient.conn)
	return client.GetSelfProfile(ctx, in)
}
