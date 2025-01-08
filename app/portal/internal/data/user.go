package data

import (
	"context"
	"database/sql"

	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmoiron/sqlx"
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

func (r *userRepo) SaveUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	insertSql := `INSERT INTO users (uid, phone, nick, icon, birth, gender, sign, state, roles, props, notify_props, update_at, created_at, delete_at) 
					VALUES (:uid, :phone, :nick, :icon, :birth, :gender, :sign, :state, :roles, :props, :notify_props, :update_at, :created_at, :delete_at)`

	tx := r.data.db.MustBegin()
	ret, err := tx.NamedExec(insertSql, u)
	if err != nil {
		tx.Rollback()
		return u, err
	}
	if err := tx.Commit(); err == nil {
		return u, nil
	} else {
		u.ID, _ = ret.LastInsertId()
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
					id, uid, phone, nick, icon, birth, gender, sign, state, roles, props, notify_props, updated, created, delete_at 
				   FROM users WHERE delete_at = 0 and phone= ?`

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

func (r *userRepo) GetUserInfoByIDs(ctx context.Context, ids []int64) ([]biz.User, error) {
	users := []biz.User{}
	query, args, err := sqlx.In(`SELECT
									id, uid, phone, nickname, icon, birth, gender, sign, state, roles, props, notify_props, update_at, create_at, delete_at 
   								FROM users WHERE id in (?)`,
		ids)
	if err != nil {
		return users, err
	}
	query = r.data.db.Rebind(query)
	rows, err := r.data.db.Queryx(query, args...)

	for rows.Next() {
		user := biz.User{}
		err = rows.StructScan(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (r *userRepo) UpdateUserProfile(ctx context.Context, user *biz.User) error {
	upProfileSql := `UPDATE users set 
					nick =?,
					icon =?,
					birth =?,
					gender =?,
					sign =?
					where id = ?`
	rs, err := r.data.db.MustExec(upProfileSql, user.Nick, user.Icon, user.Birth, user.Gender, user.Sign, user.ID).RowsAffected()
	if rs == 1 {
		return nil
	} else {
		return err
	}
}
