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
	jwt  *biz.JwtUseCase

	log *log.Helper
}

func NewUserService(user *biz.UserUseCase, jwt *biz.JwtUseCase, logger log.Logger) *UserService {
	return &UserService{
		user: user,
		jwt:  jwt,
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

	if user.State == biz.USER_STATE_BLACK_USER {
		return nil, v1.ErrorBlackUser("state %d", user.State)
	}

	//注销中用户，变更状态
	if user.State == biz.USER_STATE_WAIT_CLEAN {
		err := s.user.UpdateUserStateByID(ctx, user.Id, biz.USER_STATE_NORMAL)
		return nil, v1.ErrorUnknown("err: %v", err)
	}

	tokenStr, err := s.jwt.GenerateToken(ctx, user.Id)

	if err != nil {
		return nil, err
	}

	return &pb.LoginByMobileReply{
		Token: tokenStr,
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
	s.user.Logout(ctx)
	return &pb.LogoutReply{}, nil
}
func (s *UserService) DeRegister(ctx context.Context, req *pb.DeRegisterRequest) (*pb.DeRegisterReply, error) {
	return &pb.DeRegisterReply{}, nil
}

// get user profile
func (s *UserService) GetUserProfile(ctx context.Context, in *pb.GetUserProfileRequest) (*pb.GetUserProfileReply, error) {
	return &pb.GetUserProfileReply{}, nil
}

// get self profile
func (s *UserService) GetSelfProfile(ctx context.Context, in *pb.GetSelfProfileRequest) (*pb.GetSelfProfileReply, error) {
	return &pb.GetSelfProfileReply{}, nil
}

//filter使用
func (s *UserService) VerifyToken(ctx context.Context, token string) (string, error) {
	return s.jwt.VerifyToken(ctx, token)
}
