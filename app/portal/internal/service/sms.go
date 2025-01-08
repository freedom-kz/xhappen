package service

import (
	"context"

	pb "xhappen/api/portal/v1"
	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type SMSService struct {
	pb.UnimplementedUserServer
	sms *biz.SMSUseCase

	log *log.Helper
}

func NewSMSService(sms *biz.SMSUseCase, logger log.Logger) *SMSService {
	return &SMSService{
		sms: sms,
		log: log.NewHelper(logger),
	}
}

func (s *SMSService) SendSMSCode(ctx context.Context, req *pb.SMSCodeRequest) (*pb.SMSCodeReply, error) {
	err := s.sms.SendSMSCode(ctx, req.Mobile, req.DeviceID)

	if err != nil {
		return nil, err
	}

	return &pb.SMSCodeReply{}, nil
}
