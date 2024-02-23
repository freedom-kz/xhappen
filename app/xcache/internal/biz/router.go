package biz

import (
	"context"
	v1 "xhappen/api/router/v1"

	"github.com/go-kratos/kratos/v2/log"
)

/*
	设备
		：匿名
		：用户
	可以统一按照用户ID进行分片ID（匿名用户可以把设备作为临时用户ID）
	可以完成、点对点和点对多业务

	房间:新建socket，进行管理
		：匿名
		：用户
	统一按照房间分片管理，所有同一房间的用户连接分配socket gateway地址
*/

type Router struct {
	device map[string]DeviceRouterInfo
	user   map[uint64]UserRouterInfo
	room   map[uint64]RoomRouterInfo
}

type RouterUsecase struct {
	router *Router
	log    *log.Helper
}

func NewRouterUsecase(logger log.Logger) *RouterUsecase {
	return &RouterUsecase{
		log: log.NewHelper(logger),
		router: &Router{
			device: make(map[string]DeviceRouterInfo),
			user:   make(map[uint64]UserRouterInfo),
			room:   make(map[uint64]RoomRouterInfo),
		},
	}
}

func (s *RouterUsecase) DeviceBind(ctx context.Context, in *v1.DeviceBindRequest) (bool, error) {
	return true, nil
}

type DeviceRouterInfo struct {
	v1.DeviceBindRequest
}

type UserRouterInfo struct {
}

type RoomRouterInfo struct {
}
