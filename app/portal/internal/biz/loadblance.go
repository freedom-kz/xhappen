package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type LoadBlanceUseCase struct {
	log *log.Helper
}

func NewLoadBlanceUseCase(logger log.Logger) *SMSUseCase {
	return &SMSUseCase{
		log: log.NewHelper(log.With(logger, "module", "usecase/loadblance")),
	}
}

func (useCase *LoadBlanceUseCase) DispatchByClientID(ctx context.Context, clientId string) (string, error) {
	return "host", nil
}

func (useCase *LoadBlanceUseCase) DispatchByUserID(ctx context.Context, userID string) (string, error) {
	return "host", nil
}
