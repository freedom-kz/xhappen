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
	DeviceID string `redis:"did"`
	UserID   string `redis:"uid"`
	GwAddr   string `redis:"gw"`
}

type LoadBalanceRepo interface {
	GetGatewayPublicIPs() []string
	IsAlive(addr string) bool
	SaveDispatchInfo(ctx context.Context, deviceID string, userID string, gwAddr string) error
	GetDispatchInfoByDeviceID(ctx context.Context, deviceID string) (*DispatchInfo, bool, error)
	GetDispatchInfoByUserID(ctx context.Context, uID string) (*DispatchInfo, bool, error)
}

func NewLoadBlanceUseCase(repo LoadBalanceRepo, logger log.Logger) *LoadBlanceUseCase {
	return &LoadBlanceUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/loadblance")),
	}
}

func (useCase *LoadBlanceUseCase) DispatchUserDevice(ctx context.Context, userID string, deviceID string) (string, error) {
	dispatchInfo, exist, err := useCase.repo.GetDispatchInfoByDeviceID(ctx, deviceID)
	if err != nil {
		return "", err
	}
	if exist && dispatchInfo.UserID == userID && useCase.repo.IsAlive(dispatchInfo.GwAddr) {
		return dispatchInfo.GwAddr, nil
	}

	var addr string
	dispatchInfo, exist, err = useCase.repo.GetDispatchInfoByUserID(ctx, userID)
	if err != nil {
		return "", basic.ErrorUnknown("server err:%v", err)
	}

	if exist && useCase.repo.IsAlive(dispatchInfo.GwAddr) {
		addr = dispatchInfo.GwAddr
	} else {
		addr, err = useCase.strategyRandom()
		if err != nil {
			return "", err
		}
	}

	if err = useCase.repo.SaveDispatchInfo(ctx, deviceID, userID, addr); err != nil {
		return "", err
	}

	//客户端bind关系发生变化，踢出原绑定关系下客户端
	return addr, nil
}

// 仅读取分配信息
func (useCase *LoadBlanceUseCase) GetDispatchInfo(ctx context.Context, deviceId string) (*DispatchInfo, bool, error) {
	return useCase.repo.GetDispatchInfoByDeviceID(ctx, deviceId)
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
