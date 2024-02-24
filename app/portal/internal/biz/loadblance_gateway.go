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

type LoadBalanceRepo interface {
	GetGatewayPublicIPs() []string
	IsAlive(addr string) bool
	SaveDispatchInfo(ctx context.Context, clientId string, userId string, gwAddr string) error
	GetDispatchInfo(ctx context.Context, clientId string, userId string) (string, error)
}

func NewLoadBlanceUseCase(repo LoadBalanceRepo, logger log.Logger) *LoadBlanceUseCase {
	return &LoadBlanceUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/loadblance")),
	}
}

func (useCase *LoadBlanceUseCase) DispatchByClientID(ctx context.Context, clientId string) (string, error) {
	//1. 查询分配记录
	//2.1 不存在，分配记录并返回
	//2.2 已存在，并且目标服务器存活，返回；目标服务器不可达，重新分配
	addr, err := useCase.repo.GetDispatchInfo(ctx, clientId, "")

	if err != nil {
		return addr, err
	}

	if addr != "" && useCase.repo.IsAlive(addr) {
		return addr, nil
	}

	addr, err = useCase.strategyRandom()
	if err != nil {
		return addr, err
	}

	if err := useCase.repo.SaveDispatchInfo(ctx, clientId, "", addr); err != nil {
		return addr, basic.ErrorUnknown("server err:%v", err)
	}

	return addr, nil
}

func (useCase *LoadBlanceUseCase) DispatchByUserIDWithClientId(ctx context.Context, userID string, clientId string) (string, error) {
	//1. 删除已有的客户端相关记录
	//2. 根据用户进行分配，如同客户端单独分配逻辑
	return "host", nil
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
