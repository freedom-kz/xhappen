package client

import (
	"context"
	"os"

	pb "xhappen/api/transfer/v1"
	"xhappen/app/gateway/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	srcgrpc "google.golang.org/grpc"
)

const endpoint = "discovery:///transfer"

type PassClient struct {
	conn *srcgrpc.ClientConn
}

// 采用随机调用
func NewPassClient(conf *conf.Bootstrap, logger log.Logger) (*PassClient, func(), error) {
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

	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing the grpc client resources")
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}

	passClient := &PassClient{
		conn: conn,
	}
	return passClient, cleanup, nil
}

// 设备绑定
func (passClient *PassClient) Bind(ctx context.Context, in *pb.BindRequest) (*pb.BindReply, error) {
	client := pb.NewPassClient(passClient.conn)
	return client.Bind(ctx, in)
}

// 设备验证
func (passClient *PassClient) Auth(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
	client := pb.NewPassClient(passClient.conn)
	return client.Auth(ctx, in)
}

// 上行消息
func (passClient *PassClient) Submit(ctx context.Context, in *pb.SubmitRequest) (*pb.SubmitReply, error) {
	client := pb.NewPassClient(passClient.conn)
	return client.Submit(ctx, in)
}

// 指令消息
func (passClient *PassClient) Action(ctx context.Context, in *pb.ActionRequest) (*pb.ActionReply, error) {
	client := pb.NewPassClient(passClient.conn)
	return client.Action(ctx, in)
}

// 主动退出
func (passClient *PassClient) Quit(ctx context.Context, in *pb.QuitRequest) (*pb.QuitReply, error) {
	client := pb.NewPassClient(passClient.conn)
	return client.Quit(ctx, in)
}
