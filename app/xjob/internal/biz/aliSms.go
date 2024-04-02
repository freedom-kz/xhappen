package biz

import (
	"context"
	"os"
	"xhappen/app/xjob/internal/conf"
	"xhappen/pkg/event"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/go-kratos/kratos/v2/log"
)

type AliSMSUseCase struct {
	conf   *conf.Bootstrap
	client *dysmsapi20170525.Client
	log    *log.Helper
}

func NewAliSMSUseCase(conf *conf.Bootstrap, logger log.Logger) *AliSMSUseCase {
	config := &openapi.Config{
		AccessKeyId:     tea.String(conf.Info.Sms.AccessKeyId),
		AccessKeySecret: tea.String(conf.Info.Sms.AccessKeySecret),
	}

	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		logger.Log(log.LevelError, "modle", "biz/AliSMSUseCase", "err", err)
		os.Exit(1)
	}

	return &AliSMSUseCase{
		conf:   conf,
		client: client,
		log:    log.NewHelper(log.With(logger, "module", "usecase/alisms")),
	}
}

func (useCase *AliSMSUseCase) SendSMS(ctx context.Context, data *event.SMSCode) error {
	sendReq := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &data.Mobile,
		SignName:      &useCase.conf.Info.Sms.SignName,
		TemplateCode:  &useCase.conf.Info.Sms.TemplateCode,
		TemplateParam: &data.Authcode,
	}

	sendResp, _err := useCase.client.SendSms(sendReq)
	if _err != nil {
		return _err
	}

	code := sendResp.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		useCase.log.Info("err:", tea.StringValue(sendResp.Body.Message))
		return _err
	}
	return nil
}
