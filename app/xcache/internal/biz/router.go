package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Router struct {
	device map[string]DeviceRouterInfo 
	user   map[uint64]UserRouterInfo
	room   map[uint64]RoomRouterInfo
}

type RouterUsecase struct {
	log *log.Helper
}

func NewRouterUsecase(logger log.Logger) *RouterUsecase {
	return &RouterUsecase{log: log.NewHelper(logger)}
}

type DeviceRouterInfo struct {
}

type UserRouterInfo struct {
}

type RoomRouterInfo struct {
}
