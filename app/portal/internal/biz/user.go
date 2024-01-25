package biz

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	basic "xhappen/api/basic/v1"
	protocol "xhappen/api/protocol/v1"

	"xhappen/app/portal/internal/common"
	"xhappen/app/portal/internal/event"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	USER_STATE_NORMAL int = iota
	USER_STATE_WAIT_CLEAN
	USER_STATE_MUTE
	USER_STATE_BLACK_USER
)

type User struct {
	Id          int64     `db:"id"`
	UId         string    `db:"uid"`
	Phone       string    `db:"phone"`
	Nickname    string    `db:"nickname"`
	Icon        string    `db:"icon"`
	Birth       time.Time `db:"birth"`
	Gender      int       `db:"gender"`
	Sign        string    `db:"sign"`
	State       int       `db:"state"`
	Roles       string    `db:"roles"`
	Props       string    `db:"props"`
	NotifyProps string    `db:"notify_props"`
	Updated     int64     `db:"updated"`
	Created     int64     `db:"created"`
	DeleteAt    int64     `db:"delete_at"`
}

type UserRepo interface {
	GetUserByPhone(ctx context.Context, phone string) (*User, bool, error)
	SaveUser(ctx context.Context, g *User) (*User, error)
	UpdateUserStateByID(ctx context.Context, id int64, state int) (bool, error)
	SaveLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (err error)
	GetUserInfoByIDs(ctx context.Context, ids []int64) ([]User, error)
	GetAuthInfo(ctx context.Context, mobile string) (map[string]string, error)
	VerifyLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (bool, error)
	VerifyDayLimit(ctx context.Context, mobile string) (bool, error)
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
		if time.Now().Add(4 * time.Minute).Before(utils.TimeFromMillis(int64(createAt))) {
			return basic.ErrorRequestTooFast("mobile %s request too fast", mobile)
		}
	}

	ok, err := u.repo.VerifyDayLimit(ctx, mobile)
	if err != nil {
		return err
	}
	if !ok {
		return basic.ErrorSmsDayLimitExceed("mobile %s sms day limit", mobile)
	}

	//生成authcode并存储
	authCode := utils.GenerateAuthCode(6)
	u.log.Debugf("%s generate authCode %v", mobile, authCode)

	err = u.repo.SaveLoginAuthCode(ctx, mobile, clientId, authCode)
	if err != nil {
		return err
	}
	//发送至队列并返回
	return u.sendAuthCodeToKafka(ctx, mobile, authCode)
}

func (u *UserUseCase) LoginByMobile(ctx context.Context, mobile string, clienrId string, smsCode string) (*User, error) {
	//验证authcode
	authRet, err := u.repo.VerifyLoginAuthCode(ctx, mobile, clienrId, smsCode)
	if err != nil {
		return nil, err
	}

	if !authRet {
		return nil, basic.ErrorAuthCodeInvalid("auth code invalid")
	}

	//查找用户，不存在则新建返回
	user, exist, err := u.repo.GetUserByPhone(ctx, mobile)
	if err != nil {
		return nil, err
	}

	if exist {
		return user, nil
	}

	if !exist {
		now := time.Now().UnixNano()
		user.Phone = mobile
		user.UId = utils.GenerateId()
		user.Nickname = "用户" + mobile[len(mobile)-6:]
		user.Gender = 0
		user.Birth = time.Now()
		//新用户角色默认普通用户
		user.Roles = protocol.RoleType_ROLE_NORMAL.String()
		user.Created = now
		user.Updated = now
		user.DeleteAt = 0
		user, err = u.repo.SaveUser(ctx, user)
		if err != nil {
			return user, err
		} else {
			return user, nil
		}
	}

	return user, nil
}

func (u *UserUseCase) KickOff(ctx context.Context) error {
	return nil
}

func (u *UserUseCase) UpdateUserStateByID(ctx context.Context, id int64, state int) error {
	_, err := u.repo.UpdateUserStateByID(ctx, id, state)
	return err
}

func (u *UserUseCase) GetUserInfoByIDs(ctx context.Context, ids []int64) ([]User, error) {
	return u.repo.GetUserInfoByIDs(ctx, ids)
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
