package data

import (
	"context"

	"xhappen/app/xcache/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type sequenceRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.SequenceRepo {
	return &sequenceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *sequenceRepo) GetCurrentMaxSequence(ctx context.Context, id uint64) (uint64, error) {
	return 0, nil
}
