package data

import (
	"time"

	"xhappen/app/portal/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewJwtRepo, NewSMSRepo)

type Data struct {
	db  *sqlx.DB
	rdb *redis.Client

	log *log.Helper
}

func newDB(conf *conf.Bootstrap, logger log.Logger) *sqlx.DB {
	log := log.NewHelper(log.With(logger, "module", "portal/data/gorm"))

	db, err := sqlx.Connect("mysql", conf.Data.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
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
		db:  newDB(conf, logger),
		rdb: newRDB(conf, logger),
		log: loggger,
	}

	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}

	return d, cleanup, nil
}
