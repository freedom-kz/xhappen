package data

import (
	"context"
	"time"

	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type User struct {
	ID          int64     `gorm:"column:id"`
	UID         string    `gorm:"column:uid"`
	Phone       string    `gorm:"column:phone"`
	NickName    string    `gorm:"column:nick_name"`
	Icon        string    `gorm:"column:icon"`
	Birth       time.Time `gorm:"column:birth"`
	Gender      int       `gorm:"column:gender"`
	Sign        string    `gorm:"column:sign"`
	State       int       `gorm:"column:state"`
	Roles       string    `gorm:"column:roles"`
	Props       string    `gorm:"column:props"`
	NotifyProps string    `gorm:"column:notify_props"`
	CreateAt    int64     `gorm:"column:created_at"`
	UpdateAt    int64     `gorm:"column:updated_at"`
	DeleteAt    int64     `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "user"
}

func (r *userRepo) SaveUser(ctx context.Context, u *biz.User) (int64, error) {
	user := &User{
		UID:         u.UID,
		Phone:       u.Phone,
		NickName:    u.NickName,
		Icon:        u.Icon,
		Birth:       u.Birth,
		Gender:      u.Gender,
		Sign:        u.Sign,
		State:       u.State,
		Roles:       u.Roles,
		Props:       u.Props,
		NotifyProps: u.NotifyProps,
		CreateAt:    time.Now().UnixNano(),
		UpdateAt:    time.Now().UnixNano(),
		DeleteAt:    0,
	}

	result := r.data.DB(ctx).Create(user)
	return u.ID, result.Error
}

func (r *userRepo) UpdateUserStateByID(ctx context.Context, id int64, state int) (bool, error) {
	var ret bool
	var err error
	var user User

	err = r.data.db.First(&user, id).Error
	if err != nil {
		ret = false
		return ret, err
	}

	user.State = state
	err = r.data.db.Save(&user).Error
	if err != nil {
		ret = false
		return ret, err
	}

	return true, nil
}

func (r *userRepo) GetUserByPhone(ctx context.Context, phone string) (*biz.User, error) {
	user := &User{
		DeleteAt: 0,
		Phone:    phone,
	}
	err := r.data.db.First(user).Error
	if err != nil {
		return nil, err
	}

	return &biz.User{
		UID:         user.UID,
		Phone:       user.Phone,
		NickName:    user.NickName,
		Icon:        user.Icon,
		Birth:       user.Birth,
		Gender:      user.Gender,
		Sign:        user.Sign,
		State:       user.State,
		Roles:       user.Roles,
		Props:       user.Props,
		NotifyProps: user.NotifyProps,
		CreateAt:    user.CreateAt,
		UpdateAt:    user.UpdateAt,
		DeleteAt:    user.DeleteAt,
	}, err
}

func (r *userRepo) GetUserInfoByIDs(ctx context.Context, ids []int64) ([]*biz.User, error) {
	var users []User
	err := r.data.db.Find(users, ids).Error
	if err != nil {
		return nil, err
	}

	ret := make([]*biz.User, len(users))
	for _, u := range users {
		user := &biz.User{}
		user.ID = u.ID
		user.UID = u.UID
		user.Phone = u.Phone
		user.NickName = u.NickName
		user.Icon = u.Icon
		user.Birth = u.Birth
		user.Gender = u.Gender
		user.Sign = u.Sign
		user.State = u.State
		user.Roles = u.Roles
		user.Props = u.Props
		user.NotifyProps = u.NotifyProps
		user.CreateAt = u.CreateAt
		user.UpdateAt = u.UpdateAt
		user.DeleteAt = u.DeleteAt

		ret = append(ret, user)
	}
	return ret, err
}

func (r *userRepo) UpdateUserProfile(ctx context.Context, u *biz.User) error {
	var err error
	var user User

	err = r.data.db.First(&user, u.ID).Error
	if err != nil {
		return err
	}

	user.NickName = u.NickName
	user.Icon = u.Icon
	user.Birth = u.Birth
	user.Gender = u.Gender
	user.Sign = u.Sign

	err = r.data.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}
