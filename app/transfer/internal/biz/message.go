package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Session struct {
	SessionId uint64
}

type MessageUseCase struct {
	repo MessageRepo
	log  *log.Helper
}

func NewMessageUseCase(repo MessageRepo, logger log.Logger) *MessageUseCase {
	return &MessageUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/message")),
	}
}

type MessageRepo interface {
	SaveMessage(ctx context.Context) (err error)
	ListSyncSessions(ctx context.Context) (err error)
}

func (useCase *MessageUseCase) ListSyncSessions(ctx context.Context) ([]*Session, error) {
	return []*Session{}, nil
}
