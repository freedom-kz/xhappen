package service

import (
	"context"
	pb "xhappen/api/portal/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type ConfigService struct {
	pb.UnimplementedConfigServer

	log *log.Helper
}

func NewConfigService(logger log.Logger) *ConfigService {
	return &ConfigService{
		log: log.NewHelper(logger),
	}
}

func (c *ConfigService) GetBasicConfig(ctx context.Context, req *pb.GetBasicConfigRequest) (*pb.GetBasicConfigReply, error) {

	return &pb.GetBasicConfigReply{}, nil
}

func (c *ConfigService) GetSocketHostConfig(ctx context.Context, req *pb.GetSocketHostConfigRequest) (*pb.GetSocketHostConfigReply, error) {

	return &pb.GetSocketHostConfigReply{}, nil
}
