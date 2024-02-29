package cluster

import (
	"context"
	"xhappen/app/xcache/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

const (
	ENDPOINT_XCACHE_NAME = "xcache"
)

type Cluster struct {
	ctx                     context.Context
	state                   uint
	local_id                int
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
	logger log.Logger) (*Cluster, error) {
	return &Cluster{
		ctx:                     ctx,
		info:                    conf.Server.Info,
		registry:                registry,
		serveStateListen:        serveStateListen,
		clusterNodeModifyListen: clusterNodeModifyListen,
		log:                     log.NewHelper(log.With(logger, "module", "cluser/cluster_state")),
	}, nil
}

type localServeStateModifyListen interface {
	stateModify(state uint64, index uint64)
}

type clusterNodeModifyListen interface {
	update(services []*registry.ServiceInstance)
}

/*
初始化本机状态（按需加入）
配置监听
---集群初始化同步状态完成----
加载数据
---具备可服务状态---
监听处理
- 移除本机
- 加入集群
- 数据分片加载

对外需同步两个状态，一个为当前服务是否开启，另外一个是集群节点变更
*/
func (cluster *Cluster) Initialize() error {
	return nil
}

func (cluster *Cluster) tryJoinWithIndex(index uint64) (bool, error) {
	return false, nil
}

func (cluster *Cluster) watch() error {
	return nil
}
