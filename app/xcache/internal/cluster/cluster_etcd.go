package cluster

import (
	"context"
	"strconv"
	"time"
	"xhappen/app/xcache/internal/conf"
	"xhappen/pkg/etcd"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

const (
	ENDPOINT_XCACHE_NAME = "/microservices/customtype/xcache"
	SERVICE_NAME         = "xcache"
	REGISTER_LEASE_TTL   = 3 * time.Second
)

type Cluster struct {
	ctx                     context.Context
	state                   uint   //本机服务状态
	index                   int    //本机抢占索引
	local_ip                string //本机IP
	info                    *conf.Server_Info
	registry                *etcd.Registry
	serveStateListen        localServeStateModifyListen //状态变更通知
	clusterNodeModifyListen clusterNodeModifyListen     //集群节点变化通知
	log                     *log.Helper
}

func NewCluster(ctx context.Context,
	conf *conf.Bootstrap,
	registry *etcd.Registry,
	serveStateListen localServeStateModifyListen,
	clusterNodeModifyListen clusterNodeModifyListen,
	logger log.Logger,
) (*Cluster, error) {

	cluster := &Cluster{
		ctx:                     ctx,
		state:                   0,
		info:                    conf.Server.Info,
		registry:                registry,
		serveStateListen:        serveStateListen,
		clusterNodeModifyListen: clusterNodeModifyListen,
		log:                     log.NewHelper(log.With(logger, "module", "cluser/cluster_state")),
	}

	local_IP, err := utils.GetLocalIp()

	if err != nil {
		cluster.log.Fatalf("New cluster client get local ip err:%s", err)
		return nil, err
	}

	cluster.local_ip = local_IP

	//抢占进行服务，为真正服务node
	err = cluster.Initialize(ctx)
	if err != nil {
		cluster.log.Fatalf("New cluster client Initialize err:%s", err)
		return nil, err
	}
	//对本机发送当前状态
	cluster.serveStateListen.stateModify(cluster.state == 1, cluster.index)

	//监听集群节点数据，这里包含热备
	go func() {
		err = cluster.watchStart(ctx, SERVICE_NAME)
		if err != nil {
			cluster.log.Fatalf("New cluster client watch err:%s", err)
		}
	}()

	service, err := cluster.registry.GetService(ctx, SERVICE_NAME)
	if err != nil {
		return cluster, err
	}

	cluster.clusterNodeModifyListen.update(service)
	return cluster, err
}

type localServeStateModifyListen interface {
	stateModify(serving bool, index int)
}

type clusterNodeModifyListen interface {
	update(services []*registry.ServiceInstance)
}

// 按需加入集群
func (cluster *Cluster) Initialize(ctx context.Context) error {
	cap := int(cluster.info.Capacity)
	for i := 0; i < cap; i++ {
		//本机为待服务状态，需进行本机尝试注册
		if cluster.state == 0 {
			ok, err := cluster.tryJoinCluster(ctx, i)
			if err != nil {
				cluster.log.Errorf("try register node err:%s", err)
				return err
			}
			if ok {
				//服务状态变更
				cluster.state = 1
				cluster.index = i
				break
			}
		}
	}

	return nil
}

// 注册中心数据监听
func (cluster *Cluster) watchStart(ctx context.Context, key string) error {
	watchChan := cluster.registry.WatchWithPrefix(ctx, key)
	for {
		select {
		case <-watchChan:
			err := cluster.Initialize(ctx)
			if err != nil {
				cluster.log.Errorf("cluster watch event initialize instances err:%v", err)
				continue
			}
			ins, err := cluster.registry.GetService(ctx, SERVICE_NAME)
			if err != nil {
				cluster.log.Errorf("cluster watch event get instances err:%v", err)
				continue
			}
			cluster.clusterNodeModifyListen.update(ins)
		case <-ctx.Done():
			return nil
		}
	}
}

// 注册中心数据操作
func (cluster *Cluster) tryJoinCluster(ctx context.Context, index int) (bool, error) {
	indexStr := strconv.Itoa(index)
	key := ENDPOINT_XCACHE_NAME + "/" + indexStr
	ok, kac, err := cluster.registry.TryRegisterKVWithTTL(ctx, key, cluster.local_ip, REGISTER_LEASE_TTL)
	if ok {
		//注册成功对心跳进行处理
		go func() {
			for {
				select {
				case _, ok := <-kac:
					if !ok {
						cluster.log.Errorf("xcache cluster key:%s value:%s %s", key, cluster.index, err)
						continue
					}
				case <-cluster.ctx.Done():
					return
				}
			}
		}()

		return true, nil
	}

	return false, err
}
