package service

import (
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
