package biz

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	basic "xhappen/api/basic/v1"
	"xhappen/app/portal/internal/common"
	"xhappen/app/portal/internal/event"
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
	repo   UserRepo
	sender event.Sender
	log    *log.Helper
}

func NewUserUseCase(repo UserRepo, sender event.Sender, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		sender: sender,
		log:    log.NewHelper(log.With(logger, "module", "usecase/user")),
	}
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
		if utils.TimeFromMillis(int64(createAt)).After(time.Now().Add(4 * time.Minute)) {
			return basic.ErrorRequestTooFast("mobile %s request too fast", mobile)
		}
	}
	//生成authcode并存储
	authCode := utils.GenerateAuthCode(6)
	u.log.Debugf("generate authCode %v", authCode)

	err = u.repo.GenerateLoginAuthCode(ctx, mobile, clientId, authCode)
	if err != nil {
		return err
	}
	return u.sendAuthCodeToKafka(ctx, mobile, authCode)
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

func (u *UserUseCase) sendAuthCodeToKafka(ctx context.Context, mobile string, authcode string) error {
	smscodeMsg := &event.SMSCode{
		Mobile:   mobile,
		Authcode: authcode,
	}

	value, err := json.Marshal(smscodeMsg)

	if err != nil {
		return err
	}

	msg := event.NewMessage(mobile, value)

	go u.sender.Send(ctx, msg)
	return nil
}
