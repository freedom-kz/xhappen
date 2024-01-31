package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type Sequence struct {
}

// GreeterRepo is a Greater repo.
type SequenceRepo interface {
	GetCurrentMaxSequence(ctx context.Context, id uint64) (uint64, error)
}

// GreeterUsecase is a Greeter usecase.
type SequenceUsecase struct {
	repo SequenceRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewSequenceUsecase(repo SequenceRepo, logger log.Logger) *SequenceUsecase {
	return &SequenceUsecase{repo: repo, log: log.NewHelper(logger)}
}
