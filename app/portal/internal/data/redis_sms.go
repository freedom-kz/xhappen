package data

import (
	"context"
	"errors"
	"strconv"
	"time"
	"xhappen/app/portal/internal/common"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	DEFAULT_SEPARATOR     = ":"
	LOGIN_AUTHCODE_PREFIX = "authcode:"
	SMS_DAY_LIMIT_PREFIX  = "smsDayLimit:"
	SMS_DAY_LIMIT         = 10

	EXPIRE_AFTER_5_MINUTE = time.Minute * 5
	EXPIRE_AFTER_1_DAY    = time.Hour * 24
)

func (user *userRepo) SaveLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (err error) {
	key := LOGIN_AUTHCODE_PREFIX + mobile
	values := make(map[string]string)
	values[common.CLIENTID_KEY] = clientId
	values[common.SMSCODE_KEY] = smsCode

	//这里放入expire主要是担心失效时间设置错误导致的数据存在问题
	expire := int(utils.MillisFromTime(time.Now().Add(EXPIRE_AFTER_5_MINUTE)))
	values[common.EXPIRE_KEY] = strconv.Itoa(expire)
	err = user.data.rdb.HSet(ctx, key, values).Err()
	if err != nil {
		return
	}

	err = user.data.rdb.Expire(ctx, key, EXPIRE_AFTER_5_MINUTE).Err()
	return
}

// 获取smscode验证数据
func (user *userRepo) GetAuthInfo(ctx context.Context, mobile string) (map[string]string, error) {
	key := LOGIN_AUTHCODE_PREFIX + mobile
	kvs, err := user.data.rdb.HGetAll(ctx, key).Result()
	return kvs, err
}

func (user *userRepo) VerifyLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (bool, error) {
	key := LOGIN_AUTHCODE_PREFIX + mobile

	kvs, err := user.data.rdb.HGetAll(ctx, key).Result()

	if err != nil {
		return false, err
	}

	ev := kvs[common.EXPIRE_KEY]
	if ev != "" {
		expire, err := strconv.Atoi(ev)
		if err != nil || expire < int(utils.MillisFromTime(time.Now())) {
			return false, errors.New("smsCode expire")
		}
	} else {
		return false, errors.New("smsCode expire")
	}

	if kvs[common.CLIENTID_KEY] != clientId || kvs[common.SMSCODE_KEY] != smsCode {
		return false, errors.New("smsCode not match device")
	}

	//验证成功则删除，仅一次使用
	if user.data.rdb.Del(ctx, key).Err() != nil {
		user.log.Log(log.LevelError, "msg", "redis user del err", "key", key)
	}

	return true, nil
}

// 计算每日发送数，验证是否超过上限
// 每日计数，每日数据有效期24小时
func (user *userRepo) VerifyDayLimit(ctx context.Context, mobile string) (bool, error) {
	key := SMS_DAY_LIMIT_PREFIX + mobile + DEFAULT_SEPARATOR + utils.TodayString()
	cmd := user.data.rdb.Incr(ctx, key)
	limit, err := cmd.Result()
	if err != nil {
		return false, err
	}

	user.data.rdb.Expire(ctx, key, EXPIRE_AFTER_1_DAY)

	if limit > SMS_DAY_LIMIT {
		return false, nil
	}

	return true, nil
}
