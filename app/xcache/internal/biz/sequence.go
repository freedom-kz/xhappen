package biz

import (
	"context"
	"sync"
	"time"
	"xhappen/app/xcache/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

const defaultReplicas = 50
const defaultMaxEntries = 25 * 10000
const userSequenceStep = 10000
const loadDataTimeout = 10 * time.Second

// Greeter is a Greeter model.
type Sequence struct {
	user map[uint64]*userSequenceCache
	room map[uint64]*roomSequenceCache
}

// GreeterRepo is a Greater repo.
type SequenceRepo interface {
	ReloadAllocationUserSequence(ctx context.Context, index uint64, cap uint64, startId uint64, limit uint64) ([]*UserSequence, error)
	UpdateMaxSequence(ctx context.Context, id uint64, sequence uint64, maxSequence uint64) error
	AddUserSequence(ctx context.Context, sequence uint64, max_sequence uint64) (*UserSequence, error)
}

// GreeterUsecase is a Greeter usecase.
type SequenceUsecase struct {
	cfg      *conf.Bootstrap
	sequence *Sequence
	repo     SequenceRepo
	log      *log.Helper
}

type userSequenceCache struct {
	sync.Mutex
	index           uint64
	currentSequence uint64
	maxSequence     uint64
}

type UserSequence struct {
	Id          uint64
	Sequence    uint64
	MaxSequence uint64
}

type roomSequenceCache struct {
	sync.Mutex
	id              uint64
	currentSequence uint64
	maxSequence     uint64
}

// NewGreeterUsecase new a Greeter usecase.
func NewSequenceUsecase(cfg *conf.Bootstrap, repo SequenceRepo, logger log.Logger) *SequenceUsecase {
	/*
		1. 加载持久化数据到内存（用户/房间 待分片）
			1.1 根据配置服务集群和当前服务ID，对用户数据取模查询（ID），加载数据

		2. 在内存构建所需数据结构
			2.1 内存计算，如超过最大序列号，数据库更新
		3. 提供计算对外提供服务，分流内部计算和远程调用
	*/

	sequenceUsecase := &SequenceUsecase{cfg: cfg,
		sequence: &Sequence{user: make(map[uint64]*userSequenceCache), room: make(map[uint64]*roomSequenceCache)},
		repo:     repo,
		log:      log.NewHelper(logger)}

	//加载用户数据
	ctx, _ := context.WithTimeout(context.Background(), loadDataTimeout)
	startTime := time.Now()
	startId := 0
	limit := 100
	for {
		userSequences, err := repo.ReloadAllocationUserSequence(ctx,
			uint64(cfg.Server.Info.Index), uint64(cfg.Server.Info.Capacity), uint64(startId), uint64(limit))

		if err != nil {
			log.Fatalf("load user sequence err: %v", err)
			break
		}
		if len(userSequences) == 0 {
			log.Infof("load user sequence end.take %d", time.Since(startTime).Milliseconds())
			break
		}

		for _, userSequence := range userSequences {
			sequenceUsecase.sequence.user[userSequence.Id] =
				&userSequenceCache{index: userSequence.Id, currentSequence: userSequence.Sequence, maxSequence: userSequence.MaxSequence}
		}
	}
	//加载房间数据

	return sequenceUsecase
}

func (useCase *SequenceUsecase) GetLocalSequenceByIds(ids []uint64) {
	// for _,id := range ids {

	// }

}
