package biz

import (
	"context"
	protocol "xhappen/api/protocol"

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

type RouterUsecase struct {
	log *log.Helper

	device map[string]*DeviceInfo
	user   map[uint64]*UserRouterInfo
	room   map[uint64]*RoomRouterInfo
}

type DeviceInfo struct {
	ServerID       string
	ConnectSequece uint64
	DeviceType     protocol.DeviceType
	ClientID       string
	CurVersion     int32
}

type UserRouterInfo struct {
}

type RoomRouterInfo struct {
}

func NewRouterUsecase(logger log.Logger) *RouterUsecase {
	return &RouterUsecase{
		log: log.NewHelper(logger),

		device: make(map[string]*DeviceInfo),
		user:   make(map[uint64]*UserRouterInfo),
		room:   make(map[uint64]*RoomRouterInfo),
	}
}

func (usecase *RouterUsecase) DeviceBind(ctx context.Context, deviceInfo DeviceInfo) error {
	// usecase.device
	existing, ok := usecase.device[deviceInfo.ClientID]

	if !ok {
		usecase.device[deviceInfo.ClientID] = &deviceInfo
		return nil
	}

	if existing.ServerID != deviceInfo.ServerID {

	}
	return nil
}
