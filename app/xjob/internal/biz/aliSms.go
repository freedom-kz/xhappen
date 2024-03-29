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
		client: client,
		log:    log.NewHelper(log.With(logger, "module", "usecase/alisms")),
	}
}

func (useCase *AliSMSUseCase) SendSMS(ctx context.Context, data *event.SMSCode) error {
	sendReq := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  nil,
		SignName:      nil,
		TemplateCode:  nil,
		TemplateParam: nil,
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

// // This file is auto-generated, don't edit it. Thanks.
// package main

// import (
//   "os"
//   dysmsapi  "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
//   openapi  "github.com/alibabacloud-go/darabonba-openapi/client"
//   console  "github.com/alibabacloud-go/tea-console/client"
//   env  "github.com/alibabacloud-go/darabonba-env/client"
//   util  "github.com/alibabacloud-go/tea-utils/service"
//   time  "github.com/alibabacloud-go/darabonba-time/client"
//   string_  "github.com/alibabacloud-go/darabonba-string/client"
//   "github.com/alibabacloud-go/tea/tea"
// )

// // 使用AK&SK初始化账号Client
// func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *dysmsapi.Client, _err error) {
//   config := &openapi.Config{}
//   config.AccessKeyId = accessKeyId
//   config.AccessKeySecret = accessKeySecret
//   _result = &dysmsapi.Client{}
//   _result, _err = dysmsapi.NewClient(config)
//   return _result, _err
// }

// func _main (args []*string) (_err error) {
//   client, _err := CreateClient(env.GetEnv(tea.String("ACCESS_KEY_ID")), env.GetEnv(tea.String("ACCESS_KEY_SECRET")))
//   if _err != nil {
//     return _err
//   }

//   // 1.发送短信
//   sendReq := &dysmsapi.SendSmsRequest{
//     PhoneNumbers: args[0],
//     SignName: args[1],
//     TemplateCode: args[2],
//     TemplateParam: args[3],
//   }
//   sendResp, _err := client.SendSms(sendReq)
//   if _err != nil {
//     return _err
//   }

//   code := sendResp.Body.Code
//   if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
//     console.Log(tea.String("错误信息: " + tea.StringValue(sendResp.Body.Message)))
//     return _err
//   }

//   bizId := sendResp.Body.BizId
//   // 2. 等待 10 秒后查询结果
//   _err = util.Sleep(tea.Int(10000))
//   if _err != nil {
//     return _err
//   }
//   // 3.查询结果
//   phoneNums := string_.Split(args[0], tea.String(","), tea.Int(-1))
//   for _, phoneNum := range phoneNums {
//     queryReq := &dysmsapi.QuerySendDetailsRequest{
//       PhoneNumber: util.AssertAsString(phoneNum),
//       BizId: bizId,
//       SendDate: time.Format(tea.String("yyyyMMdd")),
//       PageSize: tea.Int64(10),
//       CurrentPage: tea.Int64(1),
//     }
//     queryResp, _err := client.QuerySendDetails(queryReq)
//     if _err != nil {
//       return _err
//     }

//     dtos := queryResp.Body.SmsSendDetailDTOs.SmsSendDetailDTO
//     // 打印结果
//     for _, dto := range dtos {
//       if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("3"))) {
//         console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 发送成功，接收时间: " + tea.StringValue(dto.ReceiveDate)))
//       } else if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("2"))) {
//         console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 发送失败"))
//       } else {
//         console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 正在发送中..."))
//       }

//     }
//   }
//   return _err
// }

// func main() {
//   err := _main(tea.StringSlice(os.Args[1:]))
//   if err != nil {
//     panic(err)
//   }
// }
