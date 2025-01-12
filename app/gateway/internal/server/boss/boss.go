package boss

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	pb "xhappen/api/gateway/v1"
	"xhappen/app/gateway/internal/client"
	"xhappen/app/gateway/internal/conf"
	"xhappen/pkg/utils"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Boss struct {
	ctx                      context.Context
	ctxCancel                context.CancelFunc
	GlobalConnectionSequence uint64
	logger                   log.Logger
	errValue                 atomic.Value
	opts                     atomic.Value
	startTime                time.Time
	serverId                 string
	protoVersion             int32
	minSupportProtoVersion   int32
	tcpListener              net.Listener
	wsListener               net.Listener
	tlsConfig                *tls.Config
	hubs                     []*Hub
	passClient               *client.PassClient
	exitChan                 chan int
	waitGroup                utils.WaitGroupWrapper
	isExiting                int32
}

func NewBoss(cfg *conf.Bootstrap, logger log.Logger, passClient *client.PassClient) *Boss {

	boss := &Boss{
		GlobalConnectionSequence: 0, //设备运行时连接序列号
		startTime:                time.Now(),
		protoVersion:             cfg.Server.Info.ProtoVersion,
		minSupportProtoVersion:   cfg.Server.Info.MinSupportProtoVersion,
		logger:                   logger,
		exitChan:                 make(chan int),
		passClient:               passClient,
	}

	host, err := utils.GetLocalIp()

	if err != nil {
		boss.logger.Log(log.LevelFatal, "get local ip", err)
		os.Exit(1)
	}

	boss.serverId = host
	boss.ctx, boss.ctxCancel = context.WithCancel(context.Background())
	boss.SwapOpts(cfg.Socket)
	boss.errValue.Store(errStore{})
	boss.hubStart()

	tlsConfig, err := buildTLSConfig(cfg.Socket.Main)
	if err != nil {
		boss.logger.Log(log.LevelFatal, "buildTLSConfig", err)
		os.Exit(1)
	}

	boss.tlsConfig = tlsConfig

	if tlsConfig == nil {
		boss.tcpListener, err = net.Listen("tcp", cfg.Socket.Main.TcpAddress)
		if err != nil {
			boss.logger.Log(log.LevelFatal, "buildTcpListener", err)
			os.Exit(1)
		}

		boss.wsListener, err = net.Listen("tcp", cfg.Socket.Main.WsAddress)
		if err != nil {
			boss.logger.Log(log.LevelFatal, "buildTcpListener", err)
			os.Exit(1)
		}
	} else {
		boss.tcpListener, err = tls.Listen("tcp", cfg.Socket.Main.TcpAddress, tlsConfig)
		if err != nil {
			boss.logger.Log(log.LevelFatal, "buildTcpListener", err)
			os.Exit(1)
		}

		boss.wsListener, err = tls.Listen("tcp", cfg.Socket.Main.WsAddress, tlsConfig)
		if err != nil {
			boss.logger.Log(log.LevelFatal, "buildTcpListener", err)
			os.Exit(1)
		}
	}
	return boss
}

// 启动执行
func (boss *Boss) Start(context.Context) error {
	exitCh := make(chan error)
	var once sync.Once
	//仅接收第一个服务运行错误
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				boss.logger.Log(log.LevelFatal, "exitFunc", err)
			}
			exitCh <- err
		})
	}

	bossServer := &BossServer{boss: boss}
	//socket协程启动
	boss.waitGroup.Wrap(func() {
		exitFunc(TCPServe(boss.tcpListener, bossServer, boss.logger))
	})

	boss.waitGroup.Wrap(func() {
		exitFunc(WsServe(boss.wsListener, bossServer, boss.logger))
	})
	//阻塞等待第一个运行错误返回
	err := <-exitCh
	return err
}

// kratos退出执行
func (boss *Boss) Stop(context.Context) error {
	//退出中，直接返回
	if !atomic.CompareAndSwapInt32(&boss.isExiting, 0, 1) {
		return fmt.Errorf("boss have exited")
	}

	//关闭TCP监听器
	if boss.tcpListener != nil {
		err := boss.tcpListener.Close()
		if err != nil {
			boss.logger.Log(log.LevelError, "tcpListener.Close", err)
		}
	}

	if boss.wsListener != nil {
		err := boss.wsListener.Close()
		if err != nil {
			boss.logger.Log(log.LevelError, "wsListener.Close", err)
		}
	}
	//关闭hub
	for _, hub := range boss.hubs {
		hub.Stop()
	}

	//发送退出信号
	close(boss.exitChan)
	//等待关闭，这里仅等待socket接收协程结束
	boss.waitGroup.Wait()
	boss.logger.Log(log.LevelInfo, "stop", "success")
	//取消函数调用
	boss.ctxCancel()
	return nil
}

// 创建并开启hub
func (boss *Boss) hubStart() {
	numberOfHubs := runtime.NumCPU() * 16
	boss.logger.Log(log.LevelInfo, "hubs", numberOfHubs)

	hubs := make([]*Hub, numberOfHubs)

	for i := 0; i < numberOfHubs; i++ {
		hubs[i] = newHub(boss)
		hubs[i].index = i
		hubs[i].Start()
	}
	boss.hubs = hubs
}

func (boss *Boss) HubStop() {
	boss.logger.Log(log.LevelInfo, "msg", "stopping websocket hub connections")
	for _, hub := range boss.hubs {
		hub.Stop()
	}
}

func (boss *Boss) getHubForUserId(userID string) *Hub {
	hash := utils.Hash(userID)
	index := hash % uint64(len(boss.hubs))
	return boss.hubs[int(index)]
}

func (boss *Boss) AddConnToHub(conn *Connection) {
	hub := boss.getHubForUserId(conn.UserId)
	hub.AddConn(conn)
}

func (boss *Boss) RemoveConnFromHub(conn *Connection) {
	hub := boss.getHubForUserId(conn.UserId)
	hub.RemoveConn(conn)
}

func (boss *Boss) GetConnFromHub(clientId string) *Connection {
	for _, hub := range boss.hubs {
		conn := hub.GetConnByCid(clientId)
		if conn != nil {
			return conn
		}
	}
	return nil
}

func (boss *Boss) SendDeliverToHubConn(done chan *errors.Error, deliver *pb.DeliverRequest) {
	hub := boss.getHubForUserId(deliver.UserID)
	hub.SendDeliverToConn(done, deliver)
}

func (boss *Boss) SendSyncToHubConn(done chan *errors.Error, sync *pb.SyncRequest) {
	hub := boss.getHubForUserId(sync.UserID)
	hub.SendSyncToConn(done, sync)
}

func (boss *Boss) SendBroadcastToHubConn(done chan *errors.Error, broadcast *pb.BroadcastRequest) {
	for _, hub := range boss.hubs {
		hub.SendBroadcastToConn(done, broadcast)
	}
}

func (boss *Boss) SendActionToHubConn(done chan *errors.Error, action *pb.ActionRequest) {
	hub := boss.getHubForUserId(action.UserID)
	hub.SendActionToConn(done, action)
}

func (boss *Boss) DisconnectedConn(done chan *errors.Error, disconnectForce *pb.DisconnectForceRequest) {
	hub := boss.getHubForUserId(disconnectForce.UserID)
	hub.DisconnectedConn(done, disconnectForce)
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
		return fmt.Sprintf("!OK - %s", err)
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
		caCertFile, err := os.ReadFile(cfg.TlsRootCAFile)
		if err != nil {
			return nil, err
		}
		if !tlsCertPool.AppendCertsFromPEM(caCertFile) {
			return nil, fmt.Errorf("failed to append certificate to pool")
		}
		tlsConfig.ClientCAs = tlsCertPool
	}
	tlsConfig.BuildNameToCertificate()

	return tlsConfig, nil
}
