package data

import (
	"context"
	"errors"
	"strconv"
	"time"
	"xhappen/app/portal/internal/common"
	"xhappen/pkg/utils"

	"github.com/redis/go-redis/v9"
)

const (
	DEFAULT_SEPARATOR = ":"
	LOGIN_PREFIX      = "login:sms:"

	EXPIRE_AFTER_5_MINUTE = time.Minute * 5
	EXPIRE_AFTER_1_DAY    = time.Hour * 24
)

var (
	REDIS_KEY_NOTFOUND error = errors.New("key not found")
)

func (user *userRepo) GenerateLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (err error) {
	key := LOGIN_PREFIX + mobile
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

func (user *userRepo) GetAuthInfo(ctx context.Context, mobile string) (map[string]string, error) {
	key := LOGIN_PREFIX + mobile
	kvs, err := user.data.rdb.HGetAll(ctx, key).Result()
	return kvs, err
}

func (user *userRepo) VerifyLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (bool, error) {
	key := LOGIN_PREFIX + mobile

	kvs, err := user.data.rdb.HGetAll(ctx, key).Result()

	if err == redis.Nil {
		return false, REDIS_KEY_NOTFOUND
	}

	if err != nil {
		return false, err
	}

	ev := kvs[common.EXPIRE_KEY]
	expire, err := strconv.Atoi(ev)
	if err != nil || expire < int(utils.MillisFromTime(time.Now())) {
		return false, err
	}

	if kvs[common.CLIENTID_KEY] != clientId || kvs[common.SMSCODE_KEY] != smsCode {
		return false, errors.New("data not match")
	}

	//验证成功则删除，仅一次使用
	user.data.rdb.Del(ctx, key)

	return true, nil
}
