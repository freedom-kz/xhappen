package boss

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"xhappen/app/gateway/internal/client"
	"xhappen/app/gateway/internal/conf"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

type Boss struct {
	ctx         context.Context
	ctxCancel   context.CancelFunc
	loggger     log.Logger
	errValue    atomic.Value
	opts        atomic.Value
	startTime   time.Time
	serverId    string
	tcpListener net.Listener
	wsListener  net.Listener
	tlsConfig   *tls.Config
	hubs        []*Hub
	passClient  *client.PassClient
	exitChan    chan int
	waitGroup   utils.WaitGroupWrapper
	isExiting   int32
}

func NewBoss(cfg *conf.Socket, loggger log.Logger, passClient *client.PassClient) *Boss {
	boss := &Boss{
		startTime:  time.Now(),
		loggger:    loggger,
		exitChan:   make(chan int),
		passClient: passClient,
	}

	boss.ctx, boss.ctxCancel = context.WithCancel(context.Background())
	boss.SwapOpts(cfg)
	boss.errValue.Store(errStore{})
	boss.hubStart()

	tlsConfig, err := buildTLSConfig(cfg.Main)
	if err != nil {
		boss.loggger.Log(log.LevelFatal, "buildTLSConfig", err)
		os.Exit(1)
	}

	boss.tlsConfig = tlsConfig

	boss.tcpListener, err = net.Listen("tcp", cfg.Main.TcpAddress)
	if err != nil {
		boss.loggger.Log(log.LevelFatal, "buildTcpListener", err)
		os.Exit(1)
	}

	boss.wsListener, err = net.Listen("tcp", cfg.Main.WsAddress)
	if err != nil {
		boss.loggger.Log(log.LevelFatal, "buildTcpListener", err)
		os.Exit(1)
	}
	return boss
}

func (boss *Boss) Start(context.Context) error {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				boss.loggger.Log(log.LevelFatal, "exitFunc", err)
			}
			exitCh <- err
		})
	}

	bossServer := &BossServer{boss: boss}

	boss.waitGroup.Wrap(func() {
		exitFunc(TCPServe(boss.tcpListener, bossServer, boss.loggger))
	})

	boss.waitGroup.Wrap(func() {
		exitFunc(WsServe(boss.tcpListener, bossServer, boss.loggger))
	})

	err := <-exitCh
	return err
}

func (boss *Boss) Stop(context.Context) error {
	//退出中，直接返回
	if !atomic.CompareAndSwapInt32(&boss.isExiting, 0, 1) {
		return fmt.Errorf("boss have exited.")
	}

	//关闭TCP监听器
	if boss.tcpListener != nil {
		err := boss.tcpListener.Close()
		if err != nil {
			boss.loggger.Log(log.LevelError, "tcpListener.Close", err)
		}
	}

	for _, hub := range boss.hubs {
		hub.Stop()
	}

	//发送退出信号
	close(boss.exitChan)
	//等待关闭
	boss.waitGroup.Wait()
	boss.loggger.Log(log.LevelInfo, "stop", "success")
	//取消函数调用
	boss.ctxCancel()
	return nil
}

func (boss *Boss) hubStart() {
	numberOfHubs := runtime.NumCPU() * 4
	boss.loggger.Log(log.LevelInfo, "hubs", numberOfHubs)

	hubs := make([]*Hub, numberOfHubs)

	for i := 0; i < numberOfHubs; i++ {
		hubs[i] = newHub(boss)
		hubs[i].index = i
		hubs[i].Start()
	}
	boss.hubs = hubs
}

type errStore struct {
	err error
}

func (boss *Boss) SetHealth(err error) {
	boss.errValue.Store(errStore{err: err})
}

func (boss *Boss) IsHealthy() bool {
	return boss.GetError() == nil
}

func (boss *Boss) GetError() error {
	errValue := boss.errValue.Load()
	return errValue.(errStore).err
}

func (boss *Boss) GetHealth() string {
	err := boss.GetError()
	if err != nil {
		return fmt.Sprintf("NOK - %s", err)
	}
	return "OK"
}

func (boss *Boss) GetStartTime() time.Time {
	return boss.startTime
}

func (n *Boss) GetConfig() *conf.Socket {
	return n.opts.Load().(*conf.Socket)
}

func (n *Boss) SwapOpts(opts *conf.Socket) {
	n.opts.Store(opts)
}

func buildTLSConfig(cfg *conf.Socket_Main) (*tls.Config, error) {
	var tlsConfig *tls.Config

	if cfg.TlsKey == "" && cfg.TlsCert == "" {
		return nil, nil
	}

	tlsClientAuthPolicy := tls.VerifyClientCertIfGiven

	cert, err := tls.LoadX509KeyPair(cfg.TlsCert, cfg.TlsKey)
	if err != nil {
		return nil, err
	}
	switch cfg.TlsClientAuthPolicy {
	case "require":
		tlsClientAuthPolicy = tls.RequireAnyClientCert
	case "require-verify":
		tlsClientAuthPolicy = tls.RequireAndVerifyClientCert
	default:
		tlsClientAuthPolicy = tls.NoClientCert
	}

	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tlsClientAuthPolicy,
		MinVersion:   tls.VersionTLS10,
	}

	if cfg.TlsRootCAFile != "" {
		tlsCertPool := x509.NewCertPool()
		caCertFile, err := ioutil.ReadFile(cfg.TlsRootCAFile)
		if err != nil {
			return nil, err
		}
		if !tlsCertPool.AppendCertsFromPEM(caCertFile) {
			return nil, errors.New("failed to append certificate to pool")
		}
		tlsConfig.ClientCAs = tlsCertPool
	}

	tlsConfig.BuildNameToCertificate()

	return tlsConfig, nil
}

func (boss *Boss) GetHubForUserId(userID string) *Hub {
	hash := utils.Hash(userID)
	index := hash % uint64(len(boss.hubs))
	return boss.hubs[int(index)]
}
