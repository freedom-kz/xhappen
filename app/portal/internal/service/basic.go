package service

import (
	"context"
	pb "xhappen/api/portal/v1"
	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ConfigService struct {
	pb.UnimplementedConfigServer
	lbUseCase *biz.LoadBlanceUseCase
	log       *log.Helper
}

func NewConfigService(lbUseCase *biz.LoadBlanceUseCase, logger log.Logger) *ConfigService {
	return &ConfigService{
		lbUseCase: lbUseCase,
		log:       log.NewHelper(logger),
	}
}

/*
插入为低频

	设备在登录前会有一次根据设备进行分配
	在用户登录时会有一次根据用户进行分配
	Gateway主机死亡会触发再次进行分配

查询为高频：

	分配即查询
	每次socket登录也会进行验证

本机应用缓存：剔除
redis缓存：暂选存储
持久化存储：读不友好
*/
func (c *ConfigService) GetBasicConfig(ctx context.Context, req *pb.GetBasicConfigRequest) (*pb.GetBasicConfigReply, error) {
	//动态sockethost，无用户参数按照client分配，有用户参数按照用户分配
	// , err := GetUserID(ctx)
	// if err != nil {
	// 	userID = req.ClientId
	// }

	return &pb.GetBasicConfigReply{}, nil
}

func (c *ConfigService) GetSocketHostConfig(ctx context.Context, req *pb.GetSocketHostConfigRequest) (*pb.GetSocketHostConfigReply, error) {

	return &pb.GetSocketHostConfigReply{}, nil
}
