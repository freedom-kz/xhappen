package biz

import (
	"context"
	"time"

	basic "xhappen/api/basic/v1"
	protocol "xhappen/api/protocol/v1"

	"xhappen/pkg/event"
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
	ID          int64     `db:"id"`
	UID         string    `db:"uid"`
	Phone       string    `db:"phone"`
	Nick        string    `db:"name"`
	Icon        string    `db:"icon"`
	Birth       time.Time `db:"birth"`
	Gender      int       `db:"gender"`
	Sign        string    `db:"sign"`
	State       int       `db:"state"`
	Roles       string    `db:"roles"`
	Props       string    `db:"props"`
	NotifyProps string    `db:"notify_props"`
	UpdateAt    int64     `db:"update_at"`
	CreateAt    int64     `db:"create_at"`
	DeleteAt    int64     `db:"delete_at"`
}

type UserRepo interface {
	GetUserByPhone(ctx context.Context, phone string) (*User, bool, error)
	SaveUser(ctx context.Context, g *User) (*User, error)
	UpdateUserStateByID(ctx context.Context, id int64, state int) (bool, error)
	GetUserInfoByIDs(ctx context.Context, ids []int64) ([]User, error)
	UpdateUserProfile(ctx context.Context, user *User) error
}

type UserUseCase struct {
	userRepo UserRepo
	smsRepo  SMSRepo
	sender   event.Sender
	log      *log.Helper
}

func NewUserUseCase(userRepo UserRepo, smsRepo SMSRepo, sender event.Sender, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		smsRepo:  smsRepo,
		sender:   sender,
		log:      log.NewHelper(log.With(logger, "module", "usecase/user")),
	}
}

func (u *UserUseCase) LoginByMobile(ctx context.Context, mobile string, deviceId string, smsCode string) (*User, error) {
	//验证authcode
	authRet, err := u.smsRepo.VerifyLoginAuthCode(ctx, mobile, deviceId, smsCode)
	if err != nil {
		return nil, err
	}

	if !authRet {
		return nil, basic.ErrorAuthCodeInvalid("auth code invalid")
	}

	//查找用户，不存在则新建返回
	user, exist, err := u.userRepo.GetUserByPhone(ctx, mobile)
	if err != nil {
		return nil, err
	}

	if exist {
		return user, nil
	}

	if !exist {
		now := time.Now().UnixNano()
		user.Phone = mobile
		user.UID = utils.GenerateId()
		user.Nick = "用户" + mobile[len(mobile)-6:]
		user.Gender = 0
		user.Birth = time.Now()
		//新用户角色默认普通用户
		user.Roles = protocol.RoleType_ROLE_NORMAL.String()
		user.CreateAt = now
		user.UpdateAt = now
		user.DeleteAt = 0
		user, err = u.userRepo.SaveUser(ctx, user)
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
	_, err := u.userRepo.UpdateUserStateByID(ctx, id, state)
	return err
}

func (u *UserUseCase) GetUserInfoByIDs(ctx context.Context, ids []int64) ([]User, error) {
	return u.userRepo.GetUserInfoByIDs(ctx, ids)
}

func (u *UserUseCase) UpdateUserProfile(ctx context.Context, user *User) error {
	return u.userRepo.UpdateUserProfile(ctx, user)
}
