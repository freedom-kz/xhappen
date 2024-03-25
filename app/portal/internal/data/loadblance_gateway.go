package data

import (
	"context"
	"errors"
	"sync"
	"time"
	"xhappen/app/portal/internal/biz"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/redis/go-redis/v9"
)

const (
	SERVICE_NAME_GATEWAY      = "gateway"
	LOAD_BLANCE_DEVICE_PREFIX = "gateway:deviceId"
	LOAD_BLANCE_USER_PREFIX   = "gateway:userid"
	LOAD_BALANCE_EXPIRE       = 30 * 24 * time.Hour
	METADATA_PUBLIC_IP        = "public_ip"
)

type LoadBlanceGwRepo struct {
	sync.RWMutex
	ctx              context.Context
	log              *log.Helper
	data             *Data
	discovery        registry.Discovery
	watcher          registry.Watcher
	gateway_publicIP []string
}

func NewLoadBlanceGwRepo(data *Data, discovery registry.Discovery, logger log.Logger) biz.LoadBalanceRepo {
	ctx := context.Background()
	services, err := discovery.GetService(ctx, SERVICE_NAME_GATEWAY)
	if err != nil {
		log.Fatal("msg", "failed connection to cluster: %v", err)
	}

	watcher, err := discovery.Watch(ctx, SERVICE_NAME_GATEWAY)
	if err != nil {
		log.Fatal("msg", "failed watch service[%s] in cluster: %v", SERVICE_NAME_GATEWAY, err)
	}

	repo := &LoadBlanceGwRepo{
		ctx:              ctx,
		data:             data,
		gateway_publicIP: make([]string, 0),
		discovery:        discovery,
		watcher:          watcher,
		log:              log.NewHelper(log.With(logger, "biz", "usecase/cluster")),
	}

	repo.update(services)

	go func() {
		for {
			services, err := watcher.Next()
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				repo.log.Errorf("http client watch service %v got unexpected error:=%v", SERVICE_NAME_GATEWAY, err)
				time.Sleep(time.Second)
				continue
			}
			repo.update(services)
		}
	}()

	return repo
}

func (repo *LoadBlanceGwRepo) update(services []*registry.ServiceInstance) {
	repo.Lock()
	defer repo.Unlock()

	publicIPs := make([]string, 0, len(services))
	for _, service := range services {
		if ip, ok := service.Metadata[METADATA_PUBLIC_IP]; ok {
			publicIPs = append(publicIPs, ip)
		}
	}
	repo.gateway_publicIP = publicIPs
}

func (repo *LoadBlanceGwRepo) GetGatewayPublicIPs() []string {
	repo.RLock()
	defer repo.RUnlock()
	return repo.gateway_publicIP
}

func (repo *LoadBlanceGwRepo) IsAlive(addr string) bool {
	return utils.StringInSlice(addr, repo.GetGatewayPublicIPs())
}

// 保存客户端&用户分配的网关信息
func (repo *LoadBlanceGwRepo) SaveDispatchInfo(ctx context.Context, deviceId string, userId string, gwAddr string) error {
	err := repo.data.rdb.HSet(ctx, LOAD_BLANCE_DEVICE_PREFIX+deviceId,
		biz.DispatchInfo{
			DeviceId: deviceId,
			UserId:   userId,
			GwAddr:   gwAddr,
		},
	).Err()
	if err != nil {
		return err
	}
	err = repo.data.rdb.Expire(ctx, LOAD_BLANCE_DEVICE_PREFIX+deviceId, LOAD_BALANCE_EXPIRE).Err()
	if err != nil {
		return err
	}

	err = repo.data.rdb.HSet(ctx, LOAD_BLANCE_USER_PREFIX+userId,
		biz.DispatchInfo{
			DeviceId: deviceId,
			GwAddr:   gwAddr,
		},
	).Err()
	if err != nil {
		return err
	}
	err = repo.data.rdb.Expire(ctx, LOAD_BLANCE_USER_PREFIX+userId, LOAD_BALANCE_EXPIRE).Err()
	if err != nil {
		return err
	}
	return nil
}

// 根据客户端为维度查找。
func (repo *LoadBlanceGwRepo) GetDispatchInfoByDeviceID(ctx context.Context, deviceId string) (*biz.DispatchInfo, bool, error) {
	var (
		dispatchInfo biz.DispatchInfo = biz.DispatchInfo{}
		exist        bool
		err          error
	)

	err = repo.data.rdb.HGetAll(ctx, LOAD_BLANCE_DEVICE_PREFIX+deviceId).Scan(&dispatchInfo)
	if err == redis.Nil {
		exist = false
		err = nil
	} else if err != nil {
		exist = false
	} else {
		exist = true
	}
	return &dispatchInfo, exist, err
}

// 根据客户端为维度查找。
func (repo *LoadBlanceGwRepo) GetDispatchInfoByUserID(ctx context.Context, uid string) (*biz.DispatchInfo, bool, error) {
	var (
		dispatchInfo biz.DispatchInfo = biz.DispatchInfo{}
		exist        bool
		err          error
	)

	err = repo.data.rdb.HGetAll(ctx, LOAD_BLANCE_USER_PREFIX+uid).Scan(&dispatchInfo)
	if err == redis.Nil {
		exist = false
		err = nil
	} else if err != nil {
		exist = false
	} else {
		exist = true
	}
	return &dispatchInfo, exist, err
}
