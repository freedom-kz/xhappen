package data

import (
	"context"
	"database/sql"
	"fmt"

	"xhappen/app/xcache/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type sequenceRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.SequenceRepo {
	return &sequenceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

/*
分页获取序列号相关数据
*/
func (r *sequenceRepo) ReloadAllocationUserSequence(ctx context.Context, index uint64, cap uint64, startId uint64, limit uint64) ([]*biz.UserSequence, error) {
	userSequences := []*biz.UserSequence{}
	selectSql := `SELECT
					id, sequence, max_sequence
				   FROM user_sequecne WHERE id%? = ?
				   and id >= startId limit ?`
	err := r.data.db.Select(userSequences,
		selectSql,
		cap, index)

	if err == sql.ErrNoRows {
		return userSequences, nil
	} else if err != nil {
		return userSequences, err
	}
	return userSequences, nil
}

/*
更新用户
*/
func (r *sequenceRepo) UpdateMaxSequence(ctx context.Context, id uint64, sequence uint64, maxSequence uint64) error {
	updateSql := `update user_sequecne set sequence = ?, max_sequence = ? WHERE id = ?`

	ret, err := r.data.db.Exec(updateSql, sequence, maxSequence, id)
	if err != nil {
		return err
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("user sequence data update exception, same rows;%v", affected)
	}
	return nil
}

func (r *sequenceRepo) AddUserSequence(ctx context.Context, sequence uint64, max_sequence uint64) (*biz.UserSequence, error) {
	insertSql := `insert into user_sequecne(sequence, max_sequence) values (?,?)`

	ret, err := r.data.db.Exec(insertSql, sequence, max_sequence)
	if err != nil {
		return nil, err
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected != 1 {
		return nil, fmt.Errorf("user sequence insert data exception, same rows;%v", affected)
	}
	id, _ := ret.LastInsertId()

	return &biz.UserSequence{
		Id:          uint64(id),
		Sequence:    sequence,
		MaxSequence: max_sequence,
	}, nil
}
