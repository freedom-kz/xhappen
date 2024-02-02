package service

import (
	"context"
	"strings"

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

// token换取用户信息
func (s *UserService) TokenAuth(ctx context.Context, req *pb.TokenAuthRequest) (*pb.TokenAuthReply, error) {
	//1. 验证token
	uid, err := s.VerifyToken(ctx, req.Token)

	if err != nil {
		return &pb.TokenAuthReply{
			Err: &v1.ErrorTokenExpire("info %v", err).Status,
		}, nil
	} else {
		return &pb.TokenAuthReply{
			Uid: uid,
		}, nil
	}
}

// 登录手机号
func (s *UserService) LoginByMobile(ctx context.Context, req *pb.LoginByMobileRequest) (*pb.LoginByMobileReply, error) {
	//登录或者注册手机号
	user, err := s.user.LoginByMobile(ctx, req.Mobile, req.ClientId, req.SmsCode)
	if err != nil {
		return nil, err
	}

	if user.State == biz.USER_STATE_BLACK_USER {
		return nil, v1.ErrorBlackUser("state %d", user.State)
	}

	//注销中用户，变更状态
	if user.State == biz.USER_STATE_WAIT_CLEAN {
		user.State = biz.USER_STATE_NORMAL
		err := s.user.UpdateUserStateByID(ctx, user.Id, user.State)
		if err != nil {
			return nil, v1.ErrorUnknown("err: %v", err)
		}

	}
	//生成token
	tokenStr, err := s.jwt.GenerateToken(ctx, user.Id)

	if err != nil {
		return nil, err
	}

	//返回信息
	return &pb.LoginByMobileReply{
		Token: tokenStr,
		User: &v1.User{
			Id:       user.Id,
			HId:      user.UId,
			Phone:    user.Phone,
			NickName: user.Nickname,
			Birth:    timestamppb.New(user.Birth),
			Icon:     user.Icon,
			Gender:   int32(user.Gender),
			Sign:     user.Sign,
			State:    int32(user.State),
			Roles:    strings.Split(user.Roles, " "),
			Created:  user.Created,
			Updated:  user.Updated,
			DeleteAt: user.DeleteAt,
		},
	}, nil
}
func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {

	//强制离线
	_, err := GetUserID(ctx)
	if err != nil {
		return &pb.LogoutReply{}, err
	}

	err = s.user.KickOff(ctx)
	if err != nil {
		return &pb.LogoutReply{}, err
	}
	// 清空token
	token, err := GetToken(ctx)
	if err != nil {
		return &pb.LogoutReply{}, err
	}

	err = s.jwt.RemoveToken(ctx, token)
	return &pb.LogoutReply{}, err
}
func (s *UserService) DeRegister(ctx context.Context, req *pb.DeRegisterRequest) (*pb.DeRegisterReply, error) {
	//强制离线
	id, err := GetUserID(ctx)
	if err != nil {
		return &pb.DeRegisterReply{}, err
	}
	err = s.user.KickOff(ctx)
	if err != nil {
		return &pb.DeRegisterReply{}, err
	}

	// 清空token
	token, err := GetToken(ctx)
	if err != nil {
		return &pb.DeRegisterReply{}, err
	}
	err = s.jwt.RemoveToken(ctx, token)
	if err != nil {
		return &pb.DeRegisterReply{}, err
	}
	//变更用户状态
	err = s.user.UpdateUserStateByID(ctx, id, biz.USER_STATE_WAIT_CLEAN)

	return &pb.DeRegisterReply{}, err
}

// get user profile
func (s *UserService) GetUserProfile(ctx context.Context, in *pb.GetUserProfileRequest) (*pb.GetUserProfileReply, error) {
	users, err := s.user.GetUserInfoByIDs(ctx, in.Ids)
	if err != nil {
		return &pb.GetUserProfileReply{}, err
	}

	profiles := make(map[int64]*v1.UserProfile, len(users))

	for _, user := range users {
		u := &v1.UserProfile{
			Id:       user.Id,
			NickName: user.Nickname,
			Icon:     user.Icon,
			Updated:  user.Updated,
			DeleteAt: user.DeleteAt,
		}
		profiles[u.Id] = u
	}

	for _, id := range in.Ids {
		if _, ok := profiles[id]; !ok {
			profiles[id] = &v1.UserProfile{}
		}
	}

	return &pb.GetUserProfileReply{
		Users: profiles,
	}, nil
}

// get self profile
func (s *UserService) GetSelfProfile(ctx context.Context, in *pb.GetSelfProfileRequest) (*pb.GetSelfProfileReply, error) {
	id, err := GetUserID(ctx)
	if err != nil {
		return &pb.GetSelfProfileReply{}, err
	}
	users, err := s.user.GetUserInfoByIDs(ctx, []int64{id})
	if err != nil {
		return &pb.GetSelfProfileReply{}, err
	}
	if len(users) != 1 {
		return &pb.GetSelfProfileReply{}, v1.ErrorUnknown("user %v not found", id)
	}
	user := users[0]
	return &pb.GetSelfProfileReply{
		User: &v1.User{
			Id:       user.Id,
			HId:      user.UId,
			Phone:    user.Phone,
			NickName: user.Nickname,
			Birth:    timestamppb.New(user.Birth),
			Icon:     user.Icon,
			Gender:   int32(user.Gender),
			Sign:     user.Sign,
			State:    int32(user.State),
			Roles:    strings.Split(user.Roles, " "),
			Created:  user.Created,
			Updated:  user.Updated,
			DeleteAt: user.DeleteAt,
		},
	}, nil
}

// filter使用
func (s *UserService) VerifyToken(ctx context.Context, token string) (string, error) {
	return s.jwt.VerifyToken(ctx, token)
}
