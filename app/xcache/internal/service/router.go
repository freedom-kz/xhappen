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

func (s *RouterService) UserDeviceBind(ctx context.Context, in *v1.DeviceBindRequest) (*v1.DeviceBindReply, error) {
	return nil, nil
}

func (s *RouterService) UserDeviceUnBind(ctx context.Context, in *v1.DeviceUnBindRequest) (*v1.DeviceUnBindReply, error) {
	return nil, nil
}

func (s *RouterService) GetRoutersByUserIds(ctx context.Context, in *v1.RoutersByUserIdsRequest) (*v1.RoutersByUserIdsReply, error) {
	return nil, nil
}

func (s *RouterService) SaveRoomRouter(ctx context.Context, in *v1.SaveRoomRouterRequest) (*v1.SaveRoomRouterReply, error) {
	return nil, nil
}

func (s *RouterService) DeleteRoomRouter(ctx context.Context, in *v1.DeleteRoomServerRequest) (*v1.SaveRoomServerReply, error) {
	return nil, nil
}

func (s *RouterService) GetRoomRouterByID(ctx context.Context, in *v1.GetRoomRouterByIDRequest) (*v1.GetRoomRouterByIDReply, error) {
	return nil, nil
}
