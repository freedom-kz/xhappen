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
	state                   uint
	index                   int
	local_ip                string
	info                    *conf.Server_Info
	registry                *etcd.Registry
	serveStateListen        localServeStateModifyListen
	clusterNodeModifyListen clusterNodeModifyListen
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
	err = cluster.watchStart()
	if err != nil {
		cluster.log.Fatalf("New cluster client watch err:%s", err)
		return nil, err
	}
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
		//需进行本机尝试注册
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

	return nil
}

// 注册中心数据监听
func (cluster *Cluster) watchStart() error {
	//本机服务状态监控及数据操作
	//所有节点数据监控，client数据操作
	return nil
}

// 注册中心数据操作
func (cluster *Cluster) tryJoinCluster(ctx context.Context, index int) (bool, error) {
	indexStr := strconv.Itoa(index)
	key := ENDPOINT_XCACHE_NAME + "/" + indexStr
	ok, err := cluster.registry.TryRegisterWithKV(ctx, key, cluster.local_ip, REGISTER_LEASE_TTL)
	return ok, err
}

func (cluster *Cluster) getXcacheClusterInfo(ctx context.Context) (map[int]string, error) {
	nodes, err := cluster.registry.GetNodeInfoWithPrefix(ctx, ENDPOINT_XCACHE_NAME)
	return nodes, err
}
