package data

import (
	"context"
	"time"
	"xhappen/app/xcache/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
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
	db  *sqlx.DB
	mdb *mongo.Client

	log *log.Helper
}

func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "portal/data"))

	d := &Data{
		db:  newDB(c, logger),
		mdb: newMDB(c, logger),
		log: helper,
	}

	mdb := d.mdb.Database(c.Data.Dms.Database)
	Collection.Message = mdb.Collection("message")

	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing the data resources")

		ctx, _ := context.WithTimeout(context.Background(), time.Second)

		if err := d.db.Close(); err != nil {
			log.Error(err)
		}

		if err := d.mdb.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}

	return d, cleanup, nil
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
