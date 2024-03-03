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
	LOAD_BLANCE_CLIENT_PREFIX = "gateway:clientid"
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

	for _, service := range services {
		if ip, ok := service.Metadata[METADATA_PUBLIC_IP]; ok {
			repo.gateway_publicIP = append(repo.gateway_publicIP, ip)
		}
	}
}

func (repo *LoadBlanceGwRepo) GetGatewayPublicIPs() []string {
	repo.RLock()
	defer repo.RUnlock()
	return repo.gateway_publicIP
}

func (repo *LoadBlanceGwRepo) IsAlive(addr string) bool {
	return utils.StringInSlice(addr, repo.gateway_publicIP)
}

func (repo *LoadBlanceGwRepo) SaveDispatchInfo(ctx context.Context, clientId string, userId string, gwAddr string) error {
	err := repo.data.rdb.Set(ctx, LOAD_BLANCE_CLIENT_PREFIX+clientId, gwAddr, LOAD_BALANCE_EXPIRE).Err()
	if err != nil {
		return err
	}
	err = repo.data.rdb.Set(ctx, LOAD_BLANCE_USER_PREFIX+userId, gwAddr, LOAD_BALANCE_EXPIRE).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *LoadBlanceGwRepo) GetDispatchInfo(ctx context.Context, clientId string, userId string) (string, error) {
	var (
		addr string
		err  error
	)

	if userId != "" {
		addr, err = repo.data.rdb.Get(ctx, LOAD_BLANCE_USER_PREFIX+userId).Result()
		if err == redis.Nil {
			addr = ""
			err = nil
		}
	} else {
		addr, err = repo.data.rdb.Get(ctx, LOAD_BLANCE_CLIENT_PREFIX+userId).Result()
		if err == redis.Nil {
			addr = ""
			err = nil
		}
	}
	return addr, err
}
