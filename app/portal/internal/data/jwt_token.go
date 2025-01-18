package data

import (
	"context"
	"time"
	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	JWT_TOKEN_PREFIX = "jwttoken:"
	EXPIRE_DAY       = time.Hour * 24
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
	err := j.data.cache.HSet(ctx, JWT_TOKEN_PREFIX+token, "id", id).Err()
	if err != nil {
		return err
	}

	if j.data.cache.Expire(ctx, JWT_TOKEN_PREFIX+token, 365*EXPIRE_DAY).Err() != nil {
		j.log.Errorf("SaveToken set expire err %s", err)
	}

	return nil
}

func (j *jwtRepo) VerifyToken(ctx context.Context, token string) (string, error) {
	cmd := j.data.cache.HGet(ctx, JWT_TOKEN_PREFIX+token, "id")
	return cmd.Result()
}

func (j *jwtRepo) RemoveToken(ctx context.Context, token string) error {
	cmd := j.data.cache.Del(ctx, JWT_TOKEN_PREFIX+token)
	return cmd.Err()
}
