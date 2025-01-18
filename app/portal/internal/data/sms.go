package data

import (
	"context"
	"errors"
	"time"
	"xhappen/app/portal/internal/biz"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

const (
	DEFAULT_SEPARATOR     = ":"
	LOGIN_AUTHCODE_PREFIX = "authcode:"
	SMS_DAY_LIMIT_PREFIX  = "smsDayLimit:"
	SMS_DAY_LIMIT         = 10

	EXPIRE_AFTER_5_MINUTE = time.Minute * 5
	EXPIRE_AFTER_1_DAY    = time.Hour * 24
)

type SMSRepo struct {
	data *Data
	log  *log.Helper
}

func NewSMSRepo(data *Data, logger log.Logger) biz.SMSRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

/*
保存手机验证码数据
当前保存map中的key有，deviceId，expire
*/
func (r *userRepo) SaveLoginAuthCode(ctx context.Context, mobile string, deviceID string, smsCode string) (err error) {
	key := LOGIN_AUTHCODE_PREFIX + mobile
	//这里放入expire主要是担心失效时间设置错误导致的数据存在问题,进行双重验证
	expire := int(utils.MillisFromTime(time.Now().Add(EXPIRE_AFTER_5_MINUTE)))
	err = r.data.cache.HSet(ctx, key, &biz.SMSInfo{
		DeviceID: deviceID,
		SMSCode:  smsCode,
		Expire:   expire,
	}).Err()
	if err != nil {
		return
	}

	err = r.data.cache.Expire(ctx, key, EXPIRE_AFTER_5_MINUTE).Err()
	return
}

// 获取smscode验证数据
func (r *userRepo) GetAuthInfo(ctx context.Context, mobile string) (*biz.SMSInfo, error) {
	var (
		smsInfo biz.SMSInfo = biz.SMSInfo{}
	)
	key := LOGIN_AUTHCODE_PREFIX + mobile
	err := r.data.cache.HGetAll(ctx, key).Scan(&smsInfo)
	if err == redis.Nil {
		err = nil
	}
	return &smsInfo, err
}

func (user *userRepo) VerifyLoginAuthCode(ctx context.Context, mobile string, deviceId string, smsCode string) (bool, error) {
	var (
		smsInfo biz.SMSInfo = biz.SMSInfo{}
	)
	key := LOGIN_AUTHCODE_PREFIX + mobile
	err := user.data.cache.HGetAll(ctx, key).Scan(&smsInfo)

	if err != nil {
		return false, err
	}

	if smsInfo.DeviceID != deviceId {
		return false, errors.New("smsCode not match device")
	}

	expire := smsInfo.Expire
	if expire < int(utils.MillisFromTime(time.Now())) {
		return false, errors.New("smsCode expire")
	}

	if smsInfo.SMSCode != smsCode {
		return false, errors.New("smsCode is invalid")
	}

	//验证成功则删除，仅一次使用
	if user.data.cache.Del(ctx, key).Err() != nil {
		user.log.Log(log.LevelError, "msg", "redis user del err", "key", key)
	}

	return true, nil
}

// 计算每日发送数，验证是否超过上限
// 每日计数，每日数据有效期24小时
func (user *userRepo) VerifyDayLimit(ctx context.Context, mobile string) (bool, error) {
	key := SMS_DAY_LIMIT_PREFIX + mobile + DEFAULT_SEPARATOR + utils.TodayString()
	cmd := user.data.cache.Incr(ctx, key)
	limit, err := cmd.Result()
	if err != nil {
		return false, err
	}

	err = user.data.cache.Expire(ctx, key, EXPIRE_AFTER_1_DAY).Err()

	if err != nil {
		return false, err
	}

	if limit > SMS_DAY_LIMIT {
		return false, nil
	}

	return true, nil
}
