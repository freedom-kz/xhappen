package etcd

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/go-kratos/kratos/v2/registry"
)

var (
	_ registry.Registrar = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)

// Option is etcd registry option.
type Option func(o *options)

type options struct {
	ctx       context.Context
	namespace string
	ttl       time.Duration
	maxRetry  int
}

// Context with registry context.
func Context(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// Namespace with registry namespace.
func Namespace(ns string) Option {
	return func(o *options) { o.namespace = ns }
}

// RegisterTTL with register ttl.
func RegisterTTL(ttl time.Duration) Option {
	return func(o *options) { o.ttl = ttl }
}

func MaxRetry(num int) Option {
	return func(o *options) { o.maxRetry = num }
}

// Registry is etcd registry.
type Registry struct {
	opts   *options
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

// New creates etcd registry
func New(client *clientv3.Client, opts ...Option) (r *Registry) {
	op := &options{
		ctx:       context.Background(),
		namespace: "/microservices",
		ttl:       time.Second * 15,
		maxRetry:  5,
	}
	for _, o := range opts {
		o(op)
	}
	return &Registry{
		opts:   op,
		client: client,
		kv:     clientv3.NewKV(client),
	}
}

// Register the registration.
func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	key := fmt.Sprintf("%s/%s/%s", r.opts.namespace, service.Name, service.ID)
	value, err := marshal(service)
	if err != nil {
		return err
	}
	if r.lease != nil {
		r.lease.Close()
	}
	r.lease = clientv3.NewLease(r.client)
	leaseID, err := r.RegisterWithKV(ctx, key, value)
	if err != nil {
		return err
	}

	go r.HeartBeat(r.opts.ctx, leaseID, key, value)
	return nil
}

// Deregister the registration.
func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	defer func() {
		if r.lease != nil {
			r.lease.Close()
		}
	}()
	key := fmt.Sprintf("%s/%s/%s", r.opts.namespace, service.Name, service.ID)
	_, err := r.client.Delete(ctx, key)
	return err
}

// GetService return the service instances in memory according to the service name.
func (r *Registry) GetService(ctx context.Context, name string) ([]*registry.ServiceInstance, error) {
	key := fmt.Sprintf("%s/%s", r.opts.namespace, name)
	resp, err := r.kv.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	items := make([]*registry.ServiceInstance, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		si, err := unmarshal(kv.Value)
		if err != nil {
			return nil, err
		}
		if si.Name != name {
			continue
		}
		items = append(items, si)
	}
	return items, nil
}

// Watch creates a watcher according to the service name.
func (r *Registry) Watch(ctx context.Context, name string) (registry.Watcher, error) {
	key := fmt.Sprintf("%s/%s", r.opts.namespace, name)
	return NewWatcher(ctx, key, name, r.client)
}

// registerWithKV create a new lease, return current leaseID
func (r *Registry) RegisterWithKV(ctx context.Context, key string, value string) (clientv3.LeaseID, error) {
	grant, err := r.lease.Grant(ctx, int64(r.opts.ttl.Seconds()))
	if err != nil {
		return 0, err
	}
	_, err = r.client.Put(ctx, key, value, clientv3.WithLease(grant.ID))
	if err != nil {
		return 0, err
	}
	return grant.ID, nil
}

// 存在则添加失败
func (r *Registry) TryJoinCluster(ctx context.Context, key string, value string, ttl time.Duration) (bool, <-chan *clientv3.LeaseKeepAliveResponse, error) {
	grant, err := r.lease.Grant(ctx, int64(r.opts.ttl.Seconds()))
	if err != nil {
		return false, nil, err
	}
	cmp := clientv3.Compare(clientv3.CreateRevision(key), "=", 0)
	put := clientv3.OpPut(key, value, clientv3.WithLease(grant.ID))
	resp, err := r.client.Txn(ctx).If(cmp).Then(put).Commit()
	if err != nil {
		return false, nil, err
	}

	if !resp.Succeeded {
		return false, nil, err
	}
	kac, err := r.client.KeepAlive(ctx, grant.ID)
	if err != nil {
		return false, nil, err
	}
	return true, kac, nil
}

func (r *Registry) GetClusterService(ctx context.Context, prefix, name string) (map[int]*registry.ServiceInstance, error) {
	//集群节点首次获取
	services, err := r.GetService(ctx, name)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s/%s", prefix, name)
	nodes, err := r.GetClusterNodes(ctx, key)
	if err != nil {
		return nil, err
	}
	ret := make(map[int]*registry.ServiceInstance)
	for idx, target := range nodes {
		for _, service := range services {
			host := service.Metadata["host"]
			if target == host {
				ret[idx] = service
				break
			}
		}
	}
	return ret, nil
}

func (r *Registry) GetClusterNodes(ctx context.Context, key string) (map[int]string, error) {
	rsp, err := r.client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	ret := make(map[int]string)
	for _, kvs := range rsp.Kvs {
		key := string(kvs.Key)
		idx := strings.Index(key, "/")
		str := key[idx:]
		index, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("GetClusterNodes strconv err:%v\n", err)
			continue
		}
		ret[index] = string(kvs.Value)
	}
	return ret, err
}

func (r *Registry) WatchWithPrefix(ctx context.Context, key string) clientv3.WatchChan {
	watcher := clientv3.NewWatcher(r.client)
	watchChan := watcher.Watch(ctx, key, clientv3.WithPrefix())
	return watchChan
}

func (r *Registry) HeartBeat(ctx context.Context, leaseID clientv3.LeaseID, key string, value string) {
	curLeaseID := leaseID
	kac, err := r.client.KeepAlive(ctx, leaseID)
	if err != nil {
		curLeaseID = 0
	}
	rand.Seed(time.Now().Unix())

	for {
		if curLeaseID == 0 {
			// try to registerWithKV
			var retreat []int
			for retryCnt := 0; retryCnt < r.opts.maxRetry; retryCnt++ {
				if ctx.Err() != nil {
					return
				}
				// prevent infinite blocking
				idChan := make(chan clientv3.LeaseID, 1)
				errChan := make(chan error, 1)
				cancelCtx, cancel := context.WithCancel(ctx)
				go func() {
					defer cancel()
					id, registerErr := r.RegisterWithKV(cancelCtx, key, value)
					if registerErr != nil {
						errChan <- registerErr
					} else {
						idChan <- id
					}
				}()

				select {
				case <-time.After(3 * time.Second):
					cancel()
					continue
				case <-errChan:
					continue
				case curLeaseID = <-idChan:
				}

				kac, err = r.client.KeepAlive(ctx, curLeaseID)
				if err == nil {
					break
				}
				retreat = append(retreat, 1<<retryCnt)
				time.Sleep(time.Duration(retreat[rand.Intn(len(retreat))]) * time.Second)
			}
			if _, ok := <-kac; !ok {
				// retry failed
				return
			}
		}

		select {
		case _, ok := <-kac:
			if !ok {
				if ctx.Err() != nil {
					// channel closed due to context cancel
					return
				}
				// need to retry registration
				curLeaseID = 0
				continue
			}
		case <-r.opts.ctx.Done():
			return
		}
	}
}

func (r *Registry) GetNodeInfoWithPrefix(ctx context.Context, prefix string) (map[int]string, error) {
	rsp, err := r.client.Get(ctx, prefix, clientv3.WithPrefix())

	if err != nil {
		return nil, err
	}

	result := make(map[int]string)
	kvs := rsp.Kvs
	for _, kv := range kvs {
		key := kv.Key
		strs := strings.Split(string(key), "/")
		if len(strs) == 0 {
			continue
		}
		index, err := strconv.Atoi(strs[len(strs)-1])

		if err != nil {
			return nil, fmt.Errorf("data[%s] is invalid", key)
		}

		val := string(kv.Value)
		result[index] = val
	}

	return result, nil
}
