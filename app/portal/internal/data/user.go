package data

import (
	"context"
	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) SaveUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ret := r.data.db.Create(u)
	if ret.Error != nil {
		return u, ret.Error
	}
	return u, nil
}

func (r *userRepo) GetUserByPhone(ctx context.Context, phone string) (*biz.User, bool, error) {
	var user = biz.User{}
	result := r.data.db.Where("phone = ?", phone).Find(&user)
	if result.Error != nil {
		return &user, false, result.Error
	}

	if result.RowsAffected == 0 {
		return &user, false, nil
	}
	return &user, true, nil
}
