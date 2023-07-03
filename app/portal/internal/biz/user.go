package biz

import (
	"context"
	"strconv"
	"time"

	basic "xhappen/api/basic/v1"
	"xhappen/app/portal/internal/common"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id          int64
	UId         string
	Roles       string
	Props       string
	NotifyProps string
	Updated     int64 `gorm:"autoUpdateTime:nano"` //更新时间
	Created     int64 `gorm:"autoCreateTime:nano"` //创建时间
	DeleteAt    int64
}

type UserRepo interface {
	Save(ctx context.Context, g *User) (*User, error)
	GenerateLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (err error)
	GetAuthInfo(ctx context.Context, mobile string) (map[string]string, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (u *UserUseCase) SendSMSCode(ctx context.Context, mobile string, clientId string) error {
	//尝试获取auth信息
	kvs, err := u.repo.GetAuthInfo(ctx, mobile)
	if err != nil {
		return err
	}
	//如果存在，验证是否发送时间在1分钟内
	if len(kvs) != 0 {
		expire := kvs[common.EXPIRE_KEY]

		createAt, err := strconv.Atoi(expire)
		if err != nil {
			return err
		}
		if utils.TimeFromMillis(int64(createAt)).Add(time.Minute).After(time.Now()) {
			return basic.ErrorRequestTooFast("mobile %s request too fast", mobile)
		}
	}
	//生成authcode并存储
	authCode := utils.GenerateAuthCode(6)
	u.log.Debugf("generate authCode %v", authCode)

	return u.repo.GenerateLoginAuthCode(ctx, mobile, clientId, authCode)
}

func (u *UserUseCase) LoginByMobile(ctx context.Context) error {
	return nil
}

func (u *UserUseCase) Logout(ctx context.Context) error {
	return nil
}

func (u *UserUseCase) Deregister(ctx context.Context) error {
	return nil
}
