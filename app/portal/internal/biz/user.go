package biz

import (
	"context"
	"time"

	basic "xhappen/api/basic/v1"
	protocol "xhappen/api/protocol/v1"

	"xhappen/pkg/event"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

const (
	USER_STATE_NORMAL int = iota
	USER_STATE_WAIT_CLEAN
	USER_STATE_MUTE
	USER_STATE_BLACK_USER
)

type User struct {
	ID          int64
	UID         string
	Phone       string
	NickName    string
	Icon        string
	Birth       time.Time
	Gender      int
	Sign        string
	State       int
	Roles       string
	Props       string
	NotifyProps string
	UpdateAt    int64
	CreateAt    int64
	DeleteAt    int64
}

type UserRepo interface {
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	SaveUser(ctx context.Context, g *User) (int64, error)
	UpdateUserStateByID(ctx context.Context, id int64, state int) (bool, error)
	GetUserInfoByIDs(ctx context.Context, ids []int64) ([]*User, error)
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
	user, err := u.userRepo.GetUserByPhone(ctx, mobile)
	if err == gorm.ErrRecordNotFound {
		user = &User{}
		now := time.Now().UnixNano()
		user.Phone = mobile
		user.UID = utils.GenerateId()
		user.NickName = "用户" + mobile[len(mobile)-6:]
		user.Gender = 0
		user.Birth = time.Now()
		//新用户角色默认普通用户
		user.Roles = protocol.RoleType_ROLE_NORMAL.String()
		user.CreateAt = now
		user.UpdateAt = now
		user.DeleteAt = 0
		id, err := u.userRepo.SaveUser(ctx, user)
		if err != nil {
			return user, err
		} else {
			user.ID = id
			return user, nil
		}
	}

	if err != nil {
		return nil, err
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

func (u *UserUseCase) GetUserInfoByIDs(ctx context.Context, ids []int64) ([]*User, error) {
	return u.userRepo.GetUserInfoByIDs(ctx, ids)
}

func (u *UserUseCase) UpdateUserProfile(ctx context.Context, user *User) error {
	return u.userRepo.UpdateUserProfile(ctx, user)
}
