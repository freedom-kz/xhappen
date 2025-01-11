package main

import (
	"flag"
	"strconv"

	"xhappen/app/gateway/internal/conf"
	"xhappen/app/gateway/internal/server/boss"
	plog "xhappen/pkg/log"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/uuid"
	etcdclient "go.etcd.io/etcd/client/v3"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = uuid.NewUUID()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	Name = "portal"
	Version = "1.0.0"
}

func newApp(bootstrap *conf.Bootstrap, logger log.Logger, gs *grpc.Server, bs *boss.Boss, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id.String()),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{
			"protoVersion":           strconv.Itoa(int(bootstrap.Server.Info.ProtoVersion)),
			"minSupportProtoVersion": strconv.Itoa(int(bootstrap.Server.Info.MinSupportProtoVersion)),
		}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			bs,
		),
		kratos.Registrar(r),
	)
}

func main() {
	flag.Parse()

	logger := log.With(newZapLog(),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	logger.Log(log.LevelDebug, "config", bc.String())

	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{bc.Data.Etcd.Addr},
	})

	if err != nil {
		log.Fatal(err)
	}

	r := etcd.New(client)

	app, cleanup, err := wireApp(&bc, r, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newZapLog() log.Logger {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "t",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	logger := plog.NewZapLogger(
		encoder,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.Development(),
	)

	return logger
}
