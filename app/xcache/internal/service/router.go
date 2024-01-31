package service

import (
	"context"

	v1 "xhappen/api/router/v1"
	"xhappen/app/xcache/internal/biz"
)

type RouterService struct {
	v1.UnimplementedRouterServer

	uc *biz.RouterUsecase
}

func NewRouterService(uc *biz.RouterUsecase) *RouterService {
	return &RouterService{uc: uc}
}

func (s *RouterService) GetServerByID(ctx context.Context, in *v1.GetServerByIDRequest) (*v1.GetServerByIDReply, error) {
	return nil, nil
}

func (s *RouterService) GetLocalServerByID(ctx context.Context, in *v1.GetServerByIDRequest) (*v1.GetServerByIDReply, error) {
	return nil, nil
}
