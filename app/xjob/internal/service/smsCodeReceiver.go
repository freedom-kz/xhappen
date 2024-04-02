package service

import (
	"context"
	"encoding/json"
	"xhappen/app/xjob/internal/biz"
	"xhappen/pkg/event"

	"github.com/go-kratos/kratos/v2/log"
)

type SmsCodeReceiverService struct {
	ctx           context.Context
	aliSMSUseCase *biz.AliSMSUseCase
	receiver      event.Receiver
	log           *log.Helper
}

func NewSmsCodeReceiverService(ctx context.Context, receiver event.Receiver, aliSMSUseCase *biz.AliSMSUseCase, logger log.Logger) *SmsCodeReceiverService {
	service := &SmsCodeReceiverService{
		ctx:           ctx,
		receiver:      receiver,
		aliSMSUseCase: aliSMSUseCase,
		log:           log.NewHelper(logger),
	}
	service.receiver.Receive(ctx, service.handler)
	return service
}

func (service *SmsCodeReceiverService) handler(ctx context.Context, msg event.Event) error {
	data := msg.Value()
	smsdataMsg := event.SMSCode{}
	err := json.Unmarshal(data, &smsdataMsg)
	if err != nil {
		return err
	}
	err = service.aliSMSUseCase.SendSMS(ctx, &smsdataMsg)
	if err != nil {
		return err
	}
	return nil
}
