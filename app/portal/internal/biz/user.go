package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id          int64
	UId         int64
	Roles       string
	Props       string
	NotifyProps string
	Updated     int64 `gorm:"autoUpdateTime:nano"` //更新时间
	Created     int64 `gorm:"autoCreateTime:nano"` //创建时间
	DeleteAt    int64
}

type UserRepo interface {
	Save(ctx context.Context, g *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}
