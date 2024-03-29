package service

import (
	"context"
	"encoding/json"
	"xhappen/pkg/event"

	"github.com/go-kratos/kratos/v2/log"
)

type SmsCodeReceiverService struct {
	ctx      context.Context
	receiver event.Receiver
	log      *log.Helper
}

func NewSmsCodeReceiverService(ctx context.Context, receiver event.Receiver, logger log.Logger) *SmsCodeReceiverService {
	service := &SmsCodeReceiverService{
		ctx:      ctx,
		receiver: receiver,
		log:      log.NewHelper(logger),
	}
	service.receiver.Receive(ctx, handler)
	return service
}

func handler(ctx context.Context, msg event.Event) error {

	data := msg.Value()
	smsdataMsg := event.SMSCode{}
	err := json.Unmarshal(data, &smsdataMsg)
	if err != nil{
		return err
	}

	

	return nil
}
