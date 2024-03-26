package service

import (
	"context"
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
	return nil
}