package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

const defaultReplicas = 50
const defaultMaxEntries = 25 * 10000
const userSequenceStep = 10000

// Greeter is a Greeter model.
type Sequence struct {
	user map[uint64]userSequenceCache
	room map[uint64]roomSequenceCache
}

// GreeterRepo is a Greater repo.
type SequenceRepo interface {
	GetCurrentMaxSequence(ctx context.Context, id uint64) (uint64, error)
}

// GreeterUsecase is a Greeter usecase.
type SequenceUsecase struct {
	repo SequenceRepo
	log  *log.Helper
}

type userSequenceCache struct {
	currentSequence uint64
	maxSequence     uint64
}

type roomSequenceCache struct {
	currentSequence uint64
	maxSequence     uint64
}

// NewGreeterUsecase new a Greeter usecase.
func NewSequenceUsecase(repo SequenceRepo, logger log.Logger) *SequenceUsecase {
	/*
		1. 加载持久化数据到内存（用户/房间 待分片）
			1.1 根据配置服务集群和当前服务ID，对用户数据取模查询（ID），加载数据

		2. 在内存构建所需数据结构
			2.1 内存计算，如超过最大序列号，数据库更新
		3. 提供计算对外提供服务，分流内部计算和远程调用
	*/

	return &SequenceUsecase{repo: repo, log: log.NewHelper(logger)}
}
