package data

import (
	"context"
	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	JWT_TOKEN_PREFIX = "jwttoken:"
)

type jwtRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewJwtRepo(data *Data, logger log.Logger) biz.JWTRepo {
	return &jwtRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (j *jwtRepo) SaveToken(ctx context.Context, token string, id string) error {
	cmd := j.data.rdb.HSet(ctx, JWT_TOKEN_PREFIX+token, "id", id)
	return cmd.Err()
}

func (j *jwtRepo) VerifyToken(ctx context.Context, token string) (string, error) {
	cmd := j.data.rdb.HGet(ctx, JWT_TOKEN_PREFIX+token, "id")
	return cmd.Result()
}

func (j *jwtRepo) RemoveToken(ctx context.Context, token string) error {
	cmd := j.data.rdb.Del(ctx, JWT_TOKEN_PREFIX+token)
	return cmd.Err()
}
