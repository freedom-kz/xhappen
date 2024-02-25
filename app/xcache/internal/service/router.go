package service

import (
	"context"

	basic "xhappen/api/basic/v1"
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

func (s *RouterService) GetServerByUserIds(ctx context.Context, in *v1.GetServerByUserIdsRequest) (*v1.GetServerByUserIdsReply, error) {
	return nil, nil
}

func (s *RouterService) GetLocalServerByUserIds(ctx context.Context, in *v1.GetLocalServerByUserIdsRequest) (*v1.GetLocalServerByUserIdsReply, error) {
	return nil, nil
}

func (s *RouterService) SaveRoomServer(ctx context.Context, in *v1.SaveRoomServerRequest) (*v1.SaveRoomServerReply, error) {
	return nil, nil
}

func (s *RouterService) SaveLocalRoomServer(ctx context.Context, in *v1.SaveLocalRoomServerRequest) (*v1.SaveLocalRoomServerReply, error) {
	return nil, nil
}

func (s *RouterService) GetRoomServerByID(ctx context.Context, in *v1.GetRoomServerByIDRequest) (*v1.GetRoomServerByIDReply, error) {
	return nil, nil
}

func (s *RouterService) GetLocalRoomServerByID(ctx context.Context, in *v1.GetLocalRoomServerByIDRequest) (*v1.GetLocalRoomServerByIDReply, error) {
	return nil, nil
}

func (s *RouterService) DeviceBind(ctx context.Context, in *v1.DeviceBindRequest) (*v1.DeviceBindReply, error) {
	/*
		1. hash计算，进行远程/本地执行
		2.远程调用，返回
		3.本地调用
		3.1 客户端不存在直接存放返回
		3.2 客户端存在，进行序列号对比，更新执行/返回错误
	*/
	request := biz.DeviceInfo{
		ClientID: in.BindInfo.ClientID,
	}
	exist, err := s.useCase.DeviceBind(ctx, &deviceBindRequest)
	if err != nil {
		return &v1.DeviceBindReply{
			Ret: false,
			Err: &basic.ErrorUnknown(err.Error()).Status,
		}, nil
	}


	return &v1.DeviceBindReply{
		Ret: true,
	}, nil
}

func (s *RouterService) DeviceUnBind(ctx context.Context, in *v1.DeviceUnBindRequest) (*v1.DeviceUnBindReply, error) {
	/*
		1. hash计算，进行远程/本地执行
		2.远程调用，返回
		3.本地调用
		3.1 客户端不存在直接返回
		3.2 客户端存在，进行序列号对比，删除或空执行返回
	*/
	return nil, nil
}

func (s *RouterService) DeviceAuth(ctx context.Context, in *v1.DeviceAuthRequest) (*v1.DeviceAuthReply, error) {
	/*
		1. hash计算，远程/本地执行
		2. 远程调用，返回
		3. 本地调用
		3.1 获取客户端bind信息

	*/
	return nil, nil
}
