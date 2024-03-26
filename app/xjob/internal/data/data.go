package data

import (
	"context"
	"time"

	"xhappen/app/xjob/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

type Data struct {
	db  *sqlx.DB
	mdb *mongo.Client
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

func newMDB(conf *conf.Bootstrap, logger log.Logger) *mongo.Client {
	opts := options.Client()
	opts.ApplyURI("mongodb://" + conf.Data.Dms.Addr)
	opts.SetMaxConnIdleTime(5 * 60 * time.Second) //空闲超时5分钟
	opts.SetMaxPoolSize(uint64(conf.Data.Dms.Idle))
	opts.SetMaxPoolSize(uint64(conf.Data.Dms.MaxConns)) //最大连接数
	//验证开关
	if conf.Data.Dms.UserName != "" {
		opts.SetAuth(options.Credential{AuthSource: "admin", Username: conf.Data.Dms.UserName, Password: conf.Data.Dms.Password})
	}
	opts.SetConnectTimeout(5 * time.Second)     //连接超时
	opts.SetReadPreference(readpref.Primary())  //只从主节点读取
	opts.SetReadConcern(readconcern.Majority()) //读策略
	journal := true
	opts.SetWriteConcern(&writeconcern.WriteConcern{
		W:       writeconcern.Majority(),
		Journal: &journal,
	})
	mongoClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("Mongo DB 初始化错误:%s\n", err.Error())
	}
	return mongoClient
}

func NewData(conf *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	loggger := log.NewHelper(log.With(logger, "module", "portal/data"))

	d := &Data{
		db:  newDB(conf, logger),
		rdb: newRDB(conf, logger),
		mdb: newMDB(conf, logger),
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
