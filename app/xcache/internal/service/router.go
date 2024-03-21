package service

import (
	"context"

	v1 "xhappen/api/router/v1"
	"xhappen/app/xcache/internal/biz"
)

type RouterService struct {
	v1.UnimplementedRouterServer

	useCase *biz.RouterUsecase
}

func NewRouterService(useCase *biz.RouterUsecase) *RouterService {
	return &RouterService{useCase: useCase}
}

func (s *RouterService) UserDeviceBind(ctx context.Context, in *v1.UserDeviceBindRequest) (*v1.UserDeviceBindReply, error) {
	return nil, nil
}

func (s *RouterService) UserDeviceUnBind(ctx context.Context, in *v1.UserDeviceUnBindRequest) (*v1.UserDeviceUnBindReply, error) {
	return nil, nil
}

func (s *RouterService) GetRoutersByUserIds(ctx context.Context, in *v1.RoutersByUserIdsRequest) (*v1.RoutersByUserIdsReply, error) {
	return nil, nil
}

func (s *RouterService) SaveRoomRouter(ctx context.Context, in *v1.RoomRouterBindRequest) (*v1.RoomRouterBindReply, error) {
	return nil, nil
}

func (s *RouterService) DeleteRoomRouter(ctx context.Context, in *v1.RoomRouterUnbindRequest) (*v1.RoomRouterUnbindReply, error) {
	return nil, nil
}

func (s *RouterService) GetRoomRouterByID(ctx context.Context, in *v1.GetRoomRouterByIDRequest) (*v1.GetRoomRouterByIDReply, error) {
	return nil, nil
}
