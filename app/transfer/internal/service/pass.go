package service

import (
	"context"
	pb_basic "xhappen/api/basic/v1"
	pb_portal "xhappen/api/portal/v1"
	pb_protocol "xhappen/api/protocol/v1"
	pb_router "xhappen/api/router/v1"
	pb_transfer "xhappen/api/transfer/v1"
	"xhappen/app/transfer/internal/biz"
	"xhappen/app/transfer/internal/client"

	"github.com/go-kratos/kratos/v2/log"
)

type PassService struct {
	pb_transfer.UnimplementedPassServer

	portalClient *client.PortalClient
	xacheClient  *client.XcahceClient
	message      *biz.MessageUseCase

	log *log.Helper
}

func NewPassService(
	portalClient *client.PortalClient,
	xcacheClient *client.XcahceClient,
	message *biz.MessageUseCase,
	logger log.Logger,
) *PassService {
	return &PassService{
		portalClient: portalClient,
		xacheClient:  xcacheClient,
		message:      message,
		log:          log.NewHelper(logger),
	}
}

func (s *PassService) Bind(ctx context.Context, in *pb_transfer.BindRequest) (*pb_transfer.BindReply, error) {
	//1. 进行终端绑定网关信息验证
	getSocketHostConfigRequest := &pb_portal.GetSocketHostConfigRequest{
		DeviceId: in.BindInfo.DeviceId,
	}

	replyHost, err := s.portalClient.GetSocketHostConfig(ctx, getSocketHostConfigRequest)

	if err != nil {
		return nil, err
	}
	if replyHost.SocketHost == "" || replyHost.SocketHost != in.ServerID {
		return &pb_transfer.BindReply{
			Ret: false,
			Err: &pb_basic.ErrorUnknown("alloc socketHost invalidate").Status,
		}, nil
	}

	return &pb_transfer.BindReply{
		Ret: true,
		Err: nil,
	}, nil
}

func (service *PassService) Auth(ctx context.Context, in *pb_transfer.AuthRequest) (*pb_transfer.AuthReply, error) {
	//1. 验证用户
	tokenAuthRequest := &pb_portal.TokenAuthRequest{
		Token:    in.AuthInfo.Token,
		DeviceId: in.DeviceId,
		RoleType: in.AuthInfo.RoleType,
	}

	authReply, err := service.portalClient.TokenAuth(ctx, tokenAuthRequest)
	if err != nil {
		return &pb_transfer.AuthReply{
			Ret: false,
			Err: &pb_basic.ErrorAuthTokenInvalid("authinfo %s.", in.AuthInfo).Status,
		}, nil
	}

	uType := pb_protocol.UserType_USER_NORMAL
	if in.AuthInfo.RoleType == pb_protocol.RoleType_ROLE_CUSTOMER_SERVICE {
		uType = pb_protocol.UserType_USER_VIRTUAL_GROUP
	}

	switch in.LoginType {
	case pb_protocol.LoginType_AUTO:
		//判断设备最后一次的bind关系，如不是当前用户，返回特定失败
	case pb_protocol.LoginType_MANUAL:
		//剔除当前设备上的用户
	}

	//2. 离线同步会话
	sessions, err := service.message.ListSyncSessions(ctx)
	if err != nil {
		return &pb_transfer.AuthReply{
			Ret:         true,
			Uid:         authReply.Uid,
			TokenExpire: authReply.TokenExpire,
			UType:       uType,
		}, nil
	}
	sids := make([]uint64, 0, len(sessions))
	for _, session := range sessions {
		sids = append(sids, session.SessionId)
	}

	//3. 路由数据存放
	deviceBindReq := &pb_router.UserDeviceBindRequest{
		UserDeviceBindInfo: &pb_router.UserDeviceBindInfo{
			DeviceID:       in.DeviceId,
			ServerID:       in.ServerID,
			ConnectSequece: in.ConnectSequece,
			CurVersion:     in.CurVersion,
			DeviceType:     in.DeviceType,
		},
	}
	deviceBindRsp, err := service.xacheClient.UserDeviceBind(ctx, deviceBindReq)
	if err != nil {
		return nil, err
	}
	if !deviceBindRsp.Ret {
		return &pb_transfer.AuthReply{
			Ret: false,
			Err: deviceBindRsp.Err,
		}, nil
	}

	return &pb_transfer.AuthReply{
		Ret:         true,
		Uid:         authReply.Uid,
		TokenExpire: authReply.TokenExpire,
		UType:       uType,
		Sessions:    sids,
	}, nil
}

func (s *PassService) Submit(ctx context.Context, in *pb_transfer.SubmitRequest) (*pb_transfer.SubmitReply, error) {
	return nil, nil
}
func (s *PassService) Action(ctx context.Context, in *pb_transfer.ActionRequest) (*pb_transfer.ActionReply, error) {
	return nil, nil
}
func (s *PassService) Quit(ctx context.Context, in *pb_transfer.QuitRequest) (*pb_transfer.QuitReply, error) {
	return nil, nil
}
