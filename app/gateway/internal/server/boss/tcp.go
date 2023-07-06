package boss

import (
	"bufio"
	"fmt"
	"net"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

func TCPServe(listener net.Listener, handler ConnHandler, logger log.Logger) error {
	logger.Log(log.LevelInfo, "tcp.listener.addr", listener.Addr())

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
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
		reader := bufio.NewReaderSize(conn, ReadBufferSize)
		writer := bufio.NewWriterSize(conn, WriteBufferSize)
		tcpConn := &TcpConn{Conn: conn, Writer: writer, Reader: reader}

		wg.Add(1)
		go func() {
			handler.Handle(tcpConn)
			wg.Done()
		}()
	}

	// wait to return until all handler goroutines complete
	wg.Wait()

	logger.Log(log.LevelInfo, "TCP: closing %s", listener.Addr())

	return nil
}

type TcpConn struct {
	net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (tc *TcpConn) Read(b []byte) (n int, err error) {
	return tc.Reader.Read(b)
}

func (tc *TcpConn) Write(b []byte) (n int, err error) {
	return tc.Writer.Write(b)
}

func (tc *TcpConn) Close() error {
	return tc.Conn.Close()
}

func (tc *TcpConn) LocalAddr() net.Addr {
	return tc.Conn.LocalAddr()
}

func (tc *TcpConn) RemoteAddr() net.Addr {
	return tc.Conn.RemoteAddr()
}

func (tc *TcpConn) SetDeadline(t time.Time) error {
	return tc.Conn.SetDeadline(t)
}
