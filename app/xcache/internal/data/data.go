package data

import (
	"context"
	"time"
	"xhappen/app/xcache/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

var (
	Collection collectionMap
)

type collectionMap struct {
	Message *mongo.Collection
}

type Data struct {
	mdb *mongo.client

	log *log.Helper
}

func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "portal/data"))

	d := &Data{
		mdb: newMDB(c, logger),
		log: helper,
	}

	mdb := d.mdb.Database(c.Data.Dms.Database)
	Collection.Message = mdb.Collection("message")

	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing the data resources")
		if err := d.mdb.Close(); err != nil {
			log.Error(err)
		}
	}

	return &Data{}, cleanup, nil
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
