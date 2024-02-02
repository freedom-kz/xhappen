package service

import (
	"context"
	basic "xhappen/api/basic/v1"
	pb_portal "xhappen/api/portal/v1"
	v1 "xhappen/api/transfer/v1"
	"xhappen/app/transfer/internal/client"

	"github.com/go-kratos/kratos/v2/log"
)

type PassService struct {
	v1.UnimplementedPassServer

	portalClient *client.PortalClient

	log *log.Helper
}

func NewPassService(portalClient *client.PortalClient, logger log.Logger) *PassService {
	return &PassService{
		portalClient: portalClient,
		log:          log.NewHelper(logger),
	}
}

func (s *PassService) Bind(ctx context.Context, in *v1.BindRequest) (*v1.BindReply, error) {

	//1. 主机验证
	getSocketHostConfigRequest := &pb_portal.GetSocketHostConfigRequest{
		ClientId: in.BindInfo.ClientID,
	}

	reply, err := s.portalClient.GetSocketHostConfig(ctx, getSocketHostConfigRequest)

	if err != nil {
		return nil, err
	}
	if reply.SocketHost == "" || reply.SocketHost != in.ServerID {
		return &v1.BindReply{
			Ret: false,
			Err: &basic.ErrorUnknown("socketHost invalidate").Status,
		}, nil
	}
	//2. 查看当前客户端的存在状态信息

	

	return &v1.BindReply{}, nil
}

func (s *PassService) Auth(ctx context.Context, in *v1.AuthRequest) (*v1.AuthReply, error) {
	// tokenAuthRequest := &pb_portal.TokenAuthRequest{
	// 	Token: in.BindInfo.,
	// }

	// reply, err := s.portalClient.TokenAuth(ctx, tokenAuthRequest)
	// if err != nil {
	// 	return &v1.BindReply{
	// 		Ret: false,
	// 		Err: &basic.ErrorAuthTokenInvalid("bindinfo %s.", in.BindInfo).Status,
	// 	}, nil
	// }
	return &v1.AuthReply{}, nil
}

func (s *PassService) Submit(ctx context.Context, in *v1.SubmitRequest) (*v1.SubmitReply, error) {
	return nil, nil
}
func (s *PassService) Action(ctx context.Context, in *v1.ActionRequest) (*v1.ActionReply, error) {
	return nil, nil
}
func (s *PassService) Quit(ctx context.Context, in *v1.QuitRequest) (*v1.QuitReply, error) {
	return nil, nil
}
