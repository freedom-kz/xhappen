package biz

import (
	"context"
	"fmt"
	"sync"
	"time"
	pb_protocol "xhappen/api/protocol/v1"

	"github.com/go-kratos/kratos/v2/log"
)

/*
	路由管理以用户为单位（这里包含匿名用户）
	关于设备为单位的管理：在用户和设备分配网关服务器时，由网关进行管理
*/

const (
	DEVICE_EVICT_INTERVAL = 300 * time.Second
)

type RouterUsecase struct {
	log *log.Helper

	sync.RWMutex
	device map[string]*DeviceRouterInfo //设备ID
	user   map[uint64]*UserRouterInfo   //用户ID
	room   map[uint64]*RoomRouterInfo   //roomID
}

type DeviceRouterInfo struct {
	IndexOfUserRouter int
	DeviceType        pb_protocol.DeviceType
	ClientID          string
	CurVersion        int32
	GatewayID         string
	GateWayVersion    int64
	ConnectSequece    uint64
	UserID            uint64
	lastActive        time.Time
}

type UserRouterInfo struct {
	devices   []*DeviceRouterInfo
	gatewayID string
}

type RoomRouterInfo struct {
	RoomID     uint64
	GatewayIDs []string
}

func NewRouterUsecase(logger log.Logger) *RouterUsecase {
	return &RouterUsecase{
		log: log.NewHelper(logger),

		device: make(map[string]*DeviceRouterInfo),
		user:   make(map[uint64]*UserRouterInfo),
		room:   make(map[uint64]*RoomRouterInfo),
	}
}

// 保存设备信息
func (usecase *RouterUsecase) SaveUserDeviceRouter(ctx context.Context, deviceInfo *DeviceRouterInfo) (*DeviceRouterInfo, error) {
	//对新设备进行判断，不允许出现不同的客户端登录在多个网关上
	usecase.RLock()
	defer usecase.RUnlock()
	userRouterInfo, ok := usecase.user[deviceInfo.UserID]
	if ok {
		if deviceInfo.GatewayID != userRouterInfo.gatewayID {
			for _, saved := range userRouterInfo.devices {
				//删除失效路由
				if time.Since(saved.lastActive) > DEVICE_EVICT_INTERVAL {
					usecase.rmDeviceInfo(saved.ClientID, saved.UserID)
					continue
				}
				if saved.ClientID != deviceInfo.ClientID {
					return nil, fmt.Errorf("add router user[%d] has diffrent gwserver", deviceInfo.UserID)
				}
			}
		}
	}
	deviceRouter := usecase.addDeviceInfo(deviceInfo)
	return deviceRouter, nil
}

// 延续设备活跃时间
func (usecase *RouterUsecase) DeviceActivate(ctx context.Context, deviceRouterInfo *DeviceRouterInfo) error {
	usecase.RLock()
	defer usecase.RUnlock()
	deviceInfo, ok := usecase.device[deviceRouterInfo.ClientID]
	if !ok {
		return fmt.Errorf("device info not exist")
	}

	if time.Since(deviceInfo.lastActive) > DEVICE_EVICT_INTERVAL {
		usecase.rmDeviceInfo(deviceInfo.ClientID, deviceInfo.UserID)
		return fmt.Errorf("device info not exist")
	}

	deviceInfo.lastActive = time.Now()
	return nil
}

// 设备解绑
func (usecase *RouterUsecase) RemoveUserDevice(ctx context.Context, deviceInfo *DeviceRouterInfo) error {
	deleted := usecase.rmDeviceInfo(deviceInfo.ClientID, deviceInfo.UserID)
	if deleted == nil {
		return nil
	}
	return fmt.Errorf("user[%d] clientid[%s] not exist", deviceInfo.UserID, deviceInfo.ClientID)
}

func (usecase *RouterUsecase) GetUserDevices(ctx context.Context, clientId string, userId uint64) ([]*DeviceRouterInfo, error) {
	usecase.RLock()
	defer usecase.RUnlock()
	if clientId == "" {
		//查找用户所有设备,查询验证和有效性
		if userRouter, ok := usecase.user[userId]; ok {
			for _, deviceRouterInfo := range userRouter.devices {
				if time.Since(deviceRouterInfo.lastActive) > DEVICE_EVICT_INTERVAL {
					usecase.rmDeviceInfo(deviceRouterInfo.ClientID, deviceRouterInfo.UserID)
				}
			}
		}

		userRouter, ok := usecase.user[userId]
		if ok {
			return userRouter.devices, nil
		} else {
			return nil, nil
		}

	} else {
		//查询指定设备
		deviceRouterInfo, ok := usecase.device[clientId]
		if !ok {
			return nil, nil
		}
		if time.Since(deviceRouterInfo.lastActive) <= DEVICE_EVICT_INTERVAL {
			//仅此为有效数据返回
			return []*DeviceRouterInfo{deviceRouterInfo}, nil
		} else {
			usecase.rmDeviceInfo(deviceRouterInfo.ClientID, deviceRouterInfo.UserID)
			return nil, nil
		}
	}
}

func (usecase *RouterUsecase) rmDeviceInfo(clientId string, userId uint64) *DeviceRouterInfo {
	usecase.Lock()
	defer usecase.Unlock()
	exist, ok := usecase.device[clientId]
	if !ok {
		usecase.log.Errorf("rm Device %s not exist", clientId)
		return nil
	}

	userRouter, ok := usecase.user[userId]
	if !ok {
		return nil
	}

	if exist.IndexOfUserRouter >= 0 && exist.IndexOfUserRouter < len(userRouter.devices) {
		copy(userRouter.devices[exist.IndexOfUserRouter:], userRouter.devices[exist.IndexOfUserRouter+1:])
		userRouter.devices = userRouter.devices[:len(userRouter.devices)-1]
	}

	if len(userRouter.devices) == 0 {
		delete(usecase.user, userId)
	}
	delete(usecase.device, clientId)
	return exist
}

// 包含新增和更新，返回已存在数据
func (usecase *RouterUsecase) addDeviceInfo(deviceInfo *DeviceRouterInfo) *DeviceRouterInfo {
	usecase.Lock()
	defer usecase.Unlock()
	//用户路由对象初始化验证
	if _, ok := usecase.user[deviceInfo.UserID]; !ok {
		usecase.user[deviceInfo.UserID] = &UserRouterInfo{
			devices:   make([]*DeviceRouterInfo, 0),
			gatewayID: deviceInfo.GatewayID,
		}

	}

	userRouter := usecase.user[deviceInfo.UserID]
	exist, ok := usecase.device[deviceInfo.ClientID]

	if !ok {
		//设备路由信息不存在，索引创建及数据填充
		deviceInfo.IndexOfUserRouter = len(userRouter.devices)
		usecase.device[deviceInfo.ClientID] = deviceInfo
		userRouter.devices = append(userRouter.devices, deviceInfo)
		return nil
	} else {
		//设备路由信息存在，复用索引，替换并返回原信息
		deviceInfo.IndexOfUserRouter = exist.IndexOfUserRouter
		usecase.device[deviceInfo.ClientID] = deviceInfo
		userRouter.devices[exist.IndexOfUserRouter] = deviceInfo
		return exist
	}
}
