package cluster

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"xhappen/app/xcache/internal/client"
	"xhappen/app/xcache/internal/conf"
	"xhappen/app/xcache/internal/service"
	"xhappen/pkg/etcd"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	ENDPOINT_CUSTOM_PREFIX = "/microservices/customtype"
	SERVICE_NAME           = "xcache"
	REGISTER_LEASE_TTL     = 3 * time.Second
)

type Cluster struct {
	ctx          context.Context
	state        uint //本机服务状态
	index        int  //本机抢占索引
	info         *conf.Server_Info
	host         string
	registry     *etcd.Registry
	xcahceClient *client.XcacheClient
	log          *log.Helper
}

/*
1. 节点状态维护（加入和离开集群）
2. 服务节点数据变更
*/
func NewCluster(ctx context.Context,
	conf *conf.Bootstrap,
	registry *etcd.Registry,
	xcahceClient *client.XcacheClient,
	logger log.Logger,
) (*Cluster, error) {

	cluster := &Cluster{
		ctx:          ctx,
		state:        0,
		info:         conf.Server.Info,
		registry:     registry,
		xcahceClient: xcahceClient,
		log:          log.NewHelper(log.With(logger, "module", "cluser/cluster_state")),
	}

	localHost, err := utils.GetLocalIp()

	if err != nil {
		cluster.log.Fatalf("New cluster client get local ip err:%s", err)
		return nil, err
	}

	cluster.host = localHost

	//抢占进行服务，为真正服务node
	err = cluster.Initialize(ctx)
	if err != nil {
		cluster.log.Fatalf("New cluster client Initialize err:%s", err)
		return nil, err
	}
	//对本机发送当前状态
	service.StateModify(cluster.state == 1, cluster.index)
	//监听集群节点数据，这里包含热备
	go func() {
		key := fmt.Sprintf("%s/%s", ENDPOINT_CUSTOM_PREFIX, SERVICE_NAME)
		err = cluster.watchStart(ctx, key)
		if err != nil {
			cluster.log.Fatalf("New cluster client watch err:%s", err)
		}
	}()

	nodes, err := cluster.registry.GetClusterService(ctx, ENDPOINT_CUSTOM_PREFIX, SERVICE_NAME)
	if err != nil {
		return cluster, err
	}

	if len(nodes) != int(conf.Server.Info.Capacity) {
		cluster.log.Infof("cluster is waiting join")
	} else {
		cluster.xcahceClient.UpdateService(nodes)
	}
	return cluster, err
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

			nodes, err := cluster.registry.GetClusterService(ctx, ENDPOINT_CUSTOM_PREFIX, SERVICE_NAME)
			if err != nil {
				cluster.log.Errorf("cluster watch event get GetClusterService err:%v", err)
				continue
			}

			if len(nodes) != int(cluster.info.Capacity) {
				cluster.log.Infof("cluster is waiting join")
			} else {
				cluster.xcahceClient.UpdateService(nodes)
			}
		case <-ctx.Done():
			return nil
		}
	}
}

// 注册中心数据操作
func (cluster *Cluster) tryJoinCluster(ctx context.Context, index int) (bool, error) {
	indexStr := strconv.Itoa(index)
	key := fmt.Sprintf("%s/%s/%s", ENDPOINT_CUSTOM_PREFIX, SERVICE_NAME, indexStr)
	ok, kac, err := cluster.registry.TryJoinCluster(ctx, key, cluster.host, REGISTER_LEASE_TTL)
	if ok {
		//注册成功对心跳进行处理
		go func() {
			for {
				select {
				case _, ok := <-kac:
					if !ok {
						cluster.log.Errorf("xcache cluster key:%s value:%s loss alive", key, cluster.index)
						//服务状态变更
						cluster.state = 0
						cluster.index = -1
						service.StateModify(cluster.state == 1, cluster.index)
						return
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
