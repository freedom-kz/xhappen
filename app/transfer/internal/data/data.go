package data

import (
	"context"
	"time"
	"xhappen/app/transfer/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var ProviderSet = wire.NewSet(NewData, NewMessageRepo)

var (
	Collection collectionMap
)

type collectionMap struct {
	Message *mongo.Collection
}

type Data struct {
	db  *mongo.client
	rdb *redis.Client

	log *log.Helper
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
	opts.SetMaxConnIdleTime(5 * 60 * time.Second)           //空闲超时5分钟
	opts.SetMaxIdleConnsPerHost(uint16(conf.Data.Dms.Idle)) //最大空闲数
	opts.SetMaxPoolSize(uint16(conf.Data.Dms.MaxConns))     //最大连接数
	//验证开关
	if conf.Data.Dms.UserName != "" {
		opts.SetAuth(options.Credential{AuthSource: "admin", Username: conf.Data.Dms.UserName, Password: conf.Data.Dms.Password})
	}
	opts.SetConnectTimeout(5 * time.Second)                                                //连接超时
	opts.SetReadPreference(readpref.Primary())                                             //只从主节点读取
	opts.SetReadConcern(readconcern.Majority())                                            //读策略
	opts.SetWriteConcern(writeconcern.New(writeconcern.WMajority(), writeconcern.J(true))) //写策略 ps: writeconcern.WTimeout(time.Second*5)写超时是否加入需测试
	mongoClient, err := mongo.Connect(context.Background(), "mongodb://"+conf.Data.Dms.Addr, opts)
	if err != nil {
		log.Fatalf("Mongo DB 初始化错误:%s\n", err.Error())
	}
	return mongoClient
}

func NewData(conf *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "portal/data"))

	d := &Data{
		db:  newMDB(conf, logger),
		rdb: newRDB(conf, logger),
		log: helper,
	}

	mdb := d.db.Database(conf.Data.Mds.Database)
	Collection.Message = mdb.Collection("message")

	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}

	return d, cleanup, nil
}
