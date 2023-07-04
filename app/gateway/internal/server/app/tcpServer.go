package app

import (
	"fmt"
	"net"
	"runtime"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
)

type TCPHandler interface {
	Handle(net.Conn)
}

func TCPServer(listener net.Listener, handler TCPHandler, logger log.Logger) error {
	logger.Log(log.LevelInfo, "listener.addr", listener.Addr())

	var wg sync.WaitGroup

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			if !strings.Contains(err.Error(), "use of closed network connection") {
				return fmt.Errorf("listener.Accept() error - %s", err)
			}

			if _, ok := err.(net.Error); ok {
				logger.Log(log.LevelWarn, "temporary Accept() failure - %s", err)
				runtime.Gosched()
				continue
			} else {
				break
			}

		}

		wg.Add(1)
		go func() {
			handler.Handle(clientConn)
			wg.Done()
		}()
	}

	// wait to return until all handler goroutines complete
	wg.Wait()

	logger.Log(log.LevelInfo, "TCP: closing %s", listener.Addr())

	return nil
}
