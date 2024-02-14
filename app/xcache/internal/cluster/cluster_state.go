package cluster

// import (
// 	"context"
// 	"xhappen/app/xcache/internal/conf"

// 	"github.com/go-kratos/kratos/v2/log"
// 	"github.com/go-kratos/kratos/v2/registry"
// )

// const (
// 	ENDPOINT_XCACHE_NAME = "xcache"
// )

// var (
// 	state      = 0  //0,running 1, working
// 	cluster_id = -1 //在working状态下的对外服务ID
// )

// type Cluster struct {
// 	ctx       context.Context
// 	info      *conf.Server_Info
// 	discovery registry.Discovery
// 	registry  registry.Registrar
// 	watcher   registry.Watcher
// 	log       *log.Helper
// }

// func NewCluster(ctx context.Context, conf *conf.Bootstrap, registrar registry.Registrar, discovery registry.Discovery, logger log.Logger) (*Cluster, error) {
// 	return &Cluster{
// 		ctx:       ctx,
// 		info:      conf.Server.Info,
// 		discovery: discovery,
// 		registry:  registrar,
// 		log:       log.NewHelper(log.With(logger, "module", "cluser/cluster_state")),
// 	}, nil
// }

// func (cluster *Cluster) initState() error {

// 	//1.尝试加入集群
// 	ins, err := cluster.discovery.GetService(cluster.ctx, ENDPOINT_XCACHE_NAME)
// 	if err != nil {
// 		return err
// 	}

// 	//2. 判断并加入
// 	for i := 0; i < int(cluster.info.Capacity); i++ {
// 		found := false

// 		for j := 0; j < len(ins); i++ {
// 			if i == j {
// 				found = true
// 				break
// 			}
// 		}

// 		if found {
// 			continue
// 		} else {
// 			ok, _ := cluster.joinWithId(i)
// 			if ok {
// 				state = 1
// 				cluster_id = i
// 			}
// 		}
// 	}
// 	//2. 配置监听
// 	watch, err := cluster.discovery.Watch(cluster.ctx, ENDPOINT_XCACHE)
// 	if err != nil {
// 		return err
// 	}

// 	cluster.watcher = watch

// 	return nil
// }

// func (cluster *Cluster) joinWithId(id int) (bool, error) {
// 	return false, nil
// }

// func (cluster *Cluster) Watch() error {
// 	_, err := cluster.discovery.Watch(cluster.ctx, ENDPOINT_XCACHE)
// 	if err != nil {
// 		return err
// 	}
// 	// for {
// 	// 	services, err := watcher.Next()
// 	// 	if err != nil {
// 	// 		done <- err
// 	// 		return
// 	// 	}
// 	// 	if r.update(services) {
// 	// 		done <- nil
// 	// 		return
// 	// 	}
// 	// }

// 	return nil
// }
