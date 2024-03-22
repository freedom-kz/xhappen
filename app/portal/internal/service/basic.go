package service

import (
	"context"
	"strconv"
	pb "xhappen/api/portal/v1"
	"xhappen/app/portal/internal/biz"
	"xhappen/app/portal/internal/conf"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

type ConfigService struct {
	pb.UnimplementedConfigServer
	lbUseCase *biz.LoadBlanceUseCase
	log       *log.Helper
	conf      *conf.Bootstrap
}

func NewConfigService(conf *conf.Bootstrap, lbUseCase *biz.LoadBlanceUseCase, logger log.Logger) *ConfigService {
	return &ConfigService{
		lbUseCase: lbUseCase,
		log:       log.NewHelper(logger),
		conf:      conf,
	}
}

// 基础数据获取，包含动态和静态配置
func (c *ConfigService) GetBasicConfig(ctx context.Context, req *pb.GetBasicConfigRequest) (*pb.GetBasicConfigReply, error) {
	//1. 获取或分配sockethost
	var (
		addr   string
		userID uint64
		err    error
	)

	if userID, err = GetUserID(ctx); err != nil {
		//无用户，生成对应设备默认匿名用户
		userID = utils.Hash(req.ClientId)
	}

	idStr := strconv.FormatUint(uint64(userID), 10)
	addr, err = c.lbUseCase.DispatchByUserIDWithClientId(ctx, req.ClientId, idStr)

	if err != nil {
		return nil, err
	} else {
		return &pb.GetBasicConfigReply{
			SocketHost:     addr,
			FileServerHost: c.conf.Info.FileServer,
		}, nil
	}
}

// 这里内部调用，仅获取socket软负载数据，不会对数据进行变更
func (c *ConfigService) GetSocketHostConfig(ctx context.Context, req *pb.GetSocketHostConfigRequest) (*pb.GetSocketHostConfigReply, error) {
	info, exist, err := c.lbUseCase.GetDispatchInfo(ctx, req.ClientId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return &pb.GetSocketHostConfigReply{}, nil
	}
	return &pb.GetSocketHostConfigReply{
		SocketHost: info.GwAddr,
	}, nil
}
