package biz

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
	basic "xhappen/api/basic/v1"
	"xhappen/app/portal/internal/common"
	"xhappen/app/portal/internal/event"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

type SMSUseCase struct {
	repo   SMSRepo
	sender event.Sender
	log    *log.Helper
}

func NewSMSUseCase(repo SMSRepo, sender event.Sender, logger log.Logger) *SMSUseCase {
	return &SMSUseCase{
		repo:   repo,
		sender: sender,
		log:    log.NewHelper(log.With(logger, "module", "usecase/sms")),
	}
}

type SMSRepo interface {
	SaveLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (err error)
	GetAuthInfo(ctx context.Context, mobile string) (map[string]string, error)
	VerifyLoginAuthCode(ctx context.Context, mobile string, clientId string, smsCode string) (bool, error)
	VerifyDayLimit(ctx context.Context, mobile string) (bool, error)
}

func (useCase *SMSUseCase) SendSMSCode(ctx context.Context, mobile string, clientId string) error {
	//尝试获取auth信息
	kvs, err := useCase.repo.GetAuthInfo(ctx, mobile)
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

	ok, err := useCase.repo.VerifyDayLimit(ctx, mobile)
	if err != nil {
		return err
	}
	if !ok {
		return basic.ErrorSmsDayLimitExceed("mobile %s sms day limit", mobile)
	}

	//生成authcode并存储
	authCode := utils.GenerateAuthCode(6)
	useCase.log.Debugf("%s generate authCode %v", mobile, authCode)

	err = useCase.repo.SaveLoginAuthCode(ctx, mobile, clientId, authCode)
	if err != nil {
		return err
	}
	//发送至队列并返回
	return useCase.sendAuthCodeToKafka(ctx, mobile, authCode)
}

func (useCase *SMSUseCase) sendAuthCodeToKafka(ctx context.Context, mobile string, authcode string) error {
	smscodeMsg := &event.SMSCode{
		Mobile:   mobile,
		Authcode: authcode,
	}

	value, err := json.Marshal(smscodeMsg)

	if err != nil {
		return err
	}

	msg := event.NewMessage(mobile, value)

	go useCase.sender.Send(ctx, msg)
	return nil
}
