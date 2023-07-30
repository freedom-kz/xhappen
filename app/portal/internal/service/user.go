package service

import (
	"context"

	v1 "xhappen/api/basic/v1"
	pb "xhappen/api/portal/v1"
	"xhappen/app/portal/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	err := s.user.SendSMSCode(ctx, req.Mobile, req.ClientId)

	if err != nil {
		return nil, err
	}

	return &pb.SMSCodeReply{}, nil
}
func (s *UserService) LoginByMobile(ctx context.Context, req *pb.LoginByMobileRequest) (*pb.LoginByMobileReply, error) {
	user, err := s.user.LoginByMobile(ctx, req.Mobile, req.ClientId, req.SmsCode)
	if err != nil {
		return nil, err
	}

	if user.State != 0 {
		return nil, v1.ErrorBlackUser("state %d", user.State)
	}

	return &pb.LoginByMobileReply{
		Token: "",
		User: &v1.User{
			Id:       uint64(user.Id),
			HId:      user.UId,
			Phone:    user.Phone,
			NickName: user.Nickname,
			Birth:    timestamppb.New(user.Birth),
			Icon:     user.Icon,
			Gender:   int32(user.Gender),
			Sign:     user.Sign,
			State:    int32(user.State),
		},
	}, nil
}
func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *UserService) DeRegister(ctx context.Context, req *pb.DeRegisterRequest) (*pb.DeRegisterReply, error) {
	return &pb.DeRegisterReply{}, nil
}
