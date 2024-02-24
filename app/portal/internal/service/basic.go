package service

import (
	"context"
	"strconv"
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
	//获取或分配sockethost
	var (
		addr string
		err  error
	)

	if userID, err := GetUserID(ctx); err == nil {
		addr, err = c.lbUseCase.DispatchByClientID(ctx, req.ClientId)
	} else {
		idStr := strconv.FormatUint(uint64(userID), 10)
		addr, err = c.lbUseCase.DispatchByUserIDWithClientId(ctx, req.ClientId, idStr)
	}

	if err != nil {
		return nil, err
	} else {
		return &pb.GetBasicConfigReply{
			SocketHost: addr,
		}, nil
	}
}

// 这里仅获取数据，不做数据变更
func (c *ConfigService) GetSocketHostConfig(ctx context.Context, req *pb.GetSocketHostConfigRequest) (*pb.GetSocketHostConfigReply, error) {
	var idStr string
	if req.UserId != 0 {
		idStr = strconv.FormatUint(uint64(req.UserId), 10)
	}
	addr, err := c.lbUseCase.GetDispatchInfo(ctx, req.ClientId, idStr)
	if err != nil {
		return nil, err
	}
	return &pb.GetSocketHostConfigReply{
		SocketHost: addr,
	}, nil
}
