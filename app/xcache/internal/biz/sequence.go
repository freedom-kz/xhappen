package biz

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"xhappen/app/xcache/internal/conf"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	DEFAULT_HASH_REPLICAS     = 50
	DEFAULT_CACHE_MAX_ENTRIES = 25 * 10000
	USER_SECTION              = 1000
	USER_SEQUENCE_STEP        = 10000
	LOAD_DATA_TIMEOUT         = 10 * time.Second
)

// GreeterUsecase is a Greeter usecase.
type SequenceUsecase struct {
	sequenceGenerater *sequenceGenerater       //内存序列号操作
	userSequence      map[uint64]*UserSequence //用户持久化数据
	roomSequence      map[uint64]*RoomSequence //房间持久化数据
	repo              SequenceRepo             //数据库操作
	log               *log.Helper              //日志接口
}

// GreeterRepo is a Greater repo.
type SequenceRepo interface {
	ReloadAllocationUserSequence(ctx context.Context, index uint64, cap uint64, startId uint64, limit uint64) ([]*UserSequence, error)
	UpdateMaxSequence(ctx context.Context, id uint64, sequence uint64, maxSequence uint64) error
	AddUserSectionSequence(ctx context.Context, sequence uint64, max_sequence uint64) (*UserSequence, error)
}

// Greeter is a Greeter model.
type sequenceGenerater struct {
	//key:userid
	user map[uint64]*userSequenceCache
	//key:roomid
	room map[uint64]*roomSequenceCache
}

type userSequenceCache struct {
	sync.RWMutex

	loadGroup       utils.Group
	useCase         *SequenceUsecase
	index           uint64
	currentSequence *uint64
	maxSequence     *uint64
}

func (cache *userSequenceCache) incrementSequence(ctx context.Context) (uint64, error) {

	// 内存数据递加
	// 判断是否为有效
	// 有效返回
	// 无效，阻塞并且单点进行持久化递进，再次进入判断循环

	next_sequence := atomic.AddUint64(cache.currentSequence, 1)
	//需要对持久化数据进行更新，并步进内存最大数据
	if next_sequence >= atomic.LoadUint64(cache.maxSequence) {
		//防击穿，相同分区索引，仅并发执行一次后返回结果
		_, err := cache.loadGroup.Do(fmt.Sprintf("%d", cache.index), func() (interface{}, error) {
			stepSequence := atomic.LoadUint64(cache.maxSequence)
			setpMaxSequence := atomic.LoadUint64(cache.maxSequence) + USER_SEQUENCE_STEP

			// update db
			err := cache.useCase.repo.UpdateMaxSequence(ctx, cache.index, stepSequence, setpMaxSequence)
			if err != nil {
				return 0, err
			} else {
				// 更新内存，最大允许序列号
				userSequence, ok := cache.useCase.userSequence[cache.index]
				if ok {
					atomic.AddUint64(userSequence.MaxSequence, USER_SEQUENCE_STEP)
				} else {
					return 0, fmt.Errorf("index %d relation data error", cache.index)
				}
				return 0, err
			}
		})

		if err != nil {
			return 0, err
		}
	}

	//后面需要注意，看是不是需要再补充一次检查执行

	return next_sequence, nil
}

func (cache *userSequenceCache) getCurrentSequence() uint64 {
	return atomic.LoadUint64(cache.currentSequence)
}

type UserSequence struct {
	Id          uint64
	Sequence    *uint64
	MaxSequence *uint64
}

type roomSequenceCache struct {
	sync.Mutex
	id              uint64
	currentSequence *uint64
	maxSequence     *uint64
}

type RoomSequence struct {
	Id          uint64
	Sequence    uint64
	MaxSequence uint64
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

	sequenceUsecase := &SequenceUsecase{
		sequenceGenerater: &sequenceGenerater{user: make(map[uint64]*userSequenceCache), room: make(map[uint64]*roomSequenceCache)},
		repo:              repo,
		log:               log.NewHelper(logger),
	}

	//加载用户数据
	ctx, _ := context.WithTimeout(context.Background(), LOAD_DATA_TIMEOUT)
	startTime := time.Now()
	startId := 0
	limit := 1000
	//加载用户序列号持久化数据，每次取1000条数据
	for {
		userSequences, err := repo.ReloadAllocationUserSequence(ctx,
			uint64(1), uint64(cfg.Server.Info.Capacity), uint64(startId), uint64(limit))

		if err != nil {
			log.Fatalf("load user sequence err: %v", err)
			break
		}
		if len(userSequences) == 0 {
			log.Infof("load user sequence end.take %d", time.Since(startTime).Milliseconds())
			break
		}

		//分段数据到内存数据的转换
		for _, userSequence := range userSequences {
			//保存用户基础数据到内存
			sequenceUsecase.userSequence[userSequence.Id] = userSequence
			for i := USER_SECTION*(userSequence.Id-1) + 1; i <= USER_SECTION*userSequence.Id; i++ {
				//复制
				sequence := *userSequence.MaxSequence
				//共同引用
				maxSequence := userSequence.MaxSequence
				sequenceUsecase.sequenceGenerater.user[i] =
					&userSequenceCache{
						useCase:         sequenceUsecase,
						index:           userSequence.Id,
						currentSequence: &sequence,
						maxSequence:     maxSequence}
			}
		}
	}
	//加载房间数据

	return sequenceUsecase
}

func (useCase *SequenceUsecase) GetLocalSequenceByIds(ctx context.Context, ids []uint64) (map[uint64]uint64, error) {
	sequences := make(map[uint64]uint64)
	for _, id := range ids {
		generater := useCase.sequenceGenerater.user[id]
		new, err := generater.incrementSequence(ctx)
		if err != nil {
			//数据库操作，更新内存数据
		} else {
			sequences[id] = new
		}
	}
	return sequences, nil
}

func (useCase *SequenceUsecase) StepUserSequence(ctx context.Context, index uint64) error {
	//持久化
	err := useCase.repo.UpdateMaxSequence(ctx, index, index, index)
	if err != nil {
		return err
	}

	//更新内存
	useCase.userSequence[index] = &UserSequence{}
	return nil
}
