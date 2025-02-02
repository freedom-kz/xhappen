package data

import (
	"context"
	"database/sql"
	"time"

	"xhappen/app/portal/internal/biz"
	"xhappen/app/portal/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewJwtRepo, NewSMSRepo, NewLoadBlanceGwRepo)

type Data struct {
	db    *gorm.DB
	cache *redis.Client
	log   *log.Helper
}

type contextTxKey struct{}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

func newDB(conf *conf.Bootstrap, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "portal/data/gorm"))

	sqlDB, err := sql.Open("mysql", conf.Data.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return gormDB
}

func newRDB(conf *conf.Bootstrap, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Data.Redis.Addr,
		Password:     conf.Data.Redis.Password,
		DB:           int(conf.Data.Redis.Db),
		DialTimeout:  conf.Data.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Data.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Data.Redis.ReadTimeout.AsDuration(),
	})

	if err := redisotel.InstrumentTracing(rdb); err != nil {
		logger.Log(log.LevelError, "msg", "redis otel init error.", "err", err)
	}
	return rdb
}

func NewData(conf *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	loggger := log.NewHelper(log.With(logger, "module", "portal/data"))

	d := &Data{
		db:    newDB(conf, logger),
		cache: newRDB(conf, logger),
		log:   loggger,
	}

	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing the data resources")
		if err := d.cache.Close(); err != nil {
			log.Error(err)
		}
	}

	return d, cleanup, nil
}
