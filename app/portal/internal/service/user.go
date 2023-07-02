package service

import (
	"context"

	pb "xhappen/api/portal/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Sendsmscode(ctx context.Context, req *pb.SMSCodeRequest) (*pb.SMSCodeReply, error) {
	return &pb.SMSCodeReply{}, nil
}
func (s *UserService) Loginbymobile(ctx context.Context, req *pb.LoginByMobileRequest) (*pb.LoginByMobileReply, error) {
	return &pb.LoginByMobileReply{}, nil
}
