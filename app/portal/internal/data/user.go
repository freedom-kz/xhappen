package data

import (
	"context"
	"database/sql"

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
	insertSql := `INSERT INTO users (uid, phone, nickname, icon, birth, gender, sign, state, roles, props, notify_props, updated, created, delete_at) 
					VALUES (:uid, :phone, :nickname, :icon, :birth, :gender, :sign, :state, :roles, :props, :notify_props, :updated, :created, :delete_at)`

	tx := r.data.db.MustBegin()
	ret, err := tx.NamedExec(insertSql, u)
	if err != nil {
		tx.Rollback()
		return u, err
	}
	if err := tx.Commit(); err == nil {
		return u, nil
	} else {
		u.Id, _ = ret.LastInsertId()
		return u, err
	}

}

func (r *userRepo) UpdateUserStateByID(ctx context.Context, id int64, state int) (bool, error) {
	upStateSql := `UPDATE users set state =? where id = ?`
	rs, err := r.data.db.MustExec(upStateSql, state, id).RowsAffected()
	if rs == 1 {
		return true, err
	} else {
		return false, err
	}
}

func (r *userRepo) GetUserByPhone(ctx context.Context, phone string) (*biz.User, bool, error) {
	user := &biz.User{}
	selectUserSql := `SELECT
					id, uid, phone, nickname, icon, birth, gender, sign, state, roles, props, notify_props, updated, created, delete_at 
				   FROM users WHERE phone= ?`

	err := r.data.db.Get(user,
		selectUserSql,
		phone)

	if err == sql.ErrNoRows {
		return user, false, nil
	} else if err != nil {
		return user, false, err
	}
	return user, true, nil
}
