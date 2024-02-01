package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type DeviceUseCase struct {
	repo DeviceRepo
	log  *log.Helper
}

func NewDeviceUseCase(repo DeviceRepo, logger log.Logger) *DeviceUseCase {
	return &DeviceUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/device")),
	}
}

type DeviceRepo interface {
	SaveDevice(ctx context.Context) (err error)
	updateDevice(ctx context.Context) (err error)
}
