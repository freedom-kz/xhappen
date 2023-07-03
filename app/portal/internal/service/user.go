package service

import (
	"context"

	pb "xhappen/api/portal/v1"
	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServer
	user *biz.UserUseCase

	log *log.Helper
}

func NewUserService(user *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		user: user,
		log:  log.NewHelper(logger),
	}
}

func (s *UserService) SendSMSCode(ctx context.Context, req *pb.SMSCodeRequest) (*pb.SMSCodeReply, error) {
	s.log.Debugf("input data %v", req)

	err := s.user.SendSMSCode(ctx, req.Mobile, req.ClientId)

	if err != nil {
		return nil, err
	}

	return &pb.SMSCodeReply{}, nil
}
func (s *UserService) LoginByMobile(ctx context.Context, req *pb.LoginByMobileRequest) (*pb.LoginByMobileReply, error) {
	return &pb.LoginByMobileReply{}, nil
}
func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *UserService) DeRegister(ctx context.Context, req *pb.DeRegisterRequest) (*pb.DeRegisterReply, error) {
	return &pb.DeRegisterReply{}, nil
}
