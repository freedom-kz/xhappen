package biz

import (
	"context"
	"math/rand"
	basic "xhappen/api/basic/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type LoadBlanceUseCase struct {
	log  *log.Helper
	repo LoadBalanceRepo
}

type DispatchInfo struct {
	ClientId string `redis:"cid"`
	UserId   string `redis:"uid"`
	GwAddr   string `redis:"gw"`
}

type LoadBalanceRepo interface {
	GetGatewayPublicIPs() []string
	IsAlive(addr string) bool
	UpsertDispatchInfo(ctx context.Context, clientId string, userId string, gwAddr string) (*DispatchInfo, error)
	GetDispatchInfo(ctx context.Context, clientId string, userId string) (*DispatchInfo, bool, error)
}

func NewLoadBlanceUseCase(repo LoadBalanceRepo, logger log.Logger) *LoadBlanceUseCase {
	return &LoadBlanceUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/loadblance")),
	}
}

func (useCase *LoadBlanceUseCase) DispatchByUserIDWithClientId(ctx context.Context, userID string, clientId string) (string, error) {
	//1. 已有分配地址，并且当前服务中，获取返回
	dispatchInfo, exist, err := useCase.repo.GetDispatchInfo(ctx, clientId, userID)
	if err != nil {
		return "", err
	}
	if exist {
		return dispatchInfo.GwAddr, nil
	}

	//2. 为用户分配地址，多个客户端登录同一网关
	addr, err := useCase.strategyRandom()
	if err != nil {
		return "", err
	}

	if _, err = useCase.repo.UpsertDispatchInfo(ctx, clientId, userID, addr); err != nil {
		//客户端原本已有网关分配绑定关系，则进行一次下线处理，新的绑定关系已形成，原客户端登录会进行校验
		return "", basic.ErrorUnknown("server err:%v", err)
	}

	return addr, nil
}

// 仅读取分配信息
func (useCase *LoadBlanceUseCase) GetDispatchInfo(ctx context.Context, clientId string, userID string) (*DispatchInfo, bool, error) {
	return useCase.repo.GetDispatchInfo(ctx, clientId, userID)
}

// 随机策略获取网关公网地址
func (useCase *LoadBlanceUseCase) strategyRandom() (string, error) {
	ins := useCase.repo.GetGatewayPublicIPs()
	if len(ins) == 0 {
		return "", basic.ErrorSerberUnavailable("all gateway is off")
	}

	idx := rand.Intn(len(ins))

	return ins[idx], nil
}
