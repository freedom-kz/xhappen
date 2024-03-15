package data

import (
	"context"
	"xhappen/app/transfer/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type messageRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageRepo(data *Data, logger log.Logger) biz.MessageRepo {
	return &messageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *messageRepo) SaveMessage(ctx context.Context) (err error) {
	return nil
}

func (repo *messageRepo) ListSyncSessions(ctx context.Context) (sessions []*biz.Session, err error) {
	return nil, nil
}
