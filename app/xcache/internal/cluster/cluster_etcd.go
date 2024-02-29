package cluster

import (
	"context"
	"xhappen/app/xcache/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	ENDPOINT_XCACHE_NAME = "xcache"
)

var (
	state      = 0  //0,running 1, working
	cluster_id = -1 //在working状态下的对外服务ID，抢占key
)

type Cluster struct {
	ctx      context.Context
	info     *conf.Server_Info
	registry *etcd.Registry
	log      *log.Helper
}

func NewCluster(ctx context.Context, conf *conf.Bootstrap, registry *etcd.Registry, logger log.Logger) (*Cluster, error) {
	return &Cluster{
		ctx:      ctx,
		info:     conf.Server.Info,
		registry: registry,
		log:      log.NewHelper(log.With(logger, "module", "cluser/cluster_state")),
	}, nil
}

/*
1. 加载注册中心数据
2. 按需加入
3. 配置监听
*/
func (cluster *Cluster) Initialize() error {
	cluster.registry.
	return nil
}

func (cluster *Cluster) TryJoin(id int) (bool, error) {
	return false, nil
}

func (cluster *Cluster) Watch() error {
	return nil
}
