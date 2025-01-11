package main

import (
	"flag"
	"os"

	"xhappen/app/portal/internal/conf"
	"xhappen/pkg/event"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	etcdclient "go.etcd.io/etcd/client/v3"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	Name = "portal"
	Version = "1.0.0"
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(r),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.Caller(3),
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

	logger.Log(log.LevelInfo, "config", bc)

	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{bc.Data.Etcd.Addr},
	})
	if err != nil {
		log.Fatal(err)
	}

	etcdclient := etcd.New(client)

	smsCodeSender, err := event.NewKafkaSender([]string{bc.Data.Kafka.Addr}, bc.Data.Kafka.SmsCodeTopic)
	if err != nil {
		log.Fatal(err)
	}

	app, cleanup, err := wireApp(&bc, logger, etcdclient, etcdclient, smsCodeSender)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
