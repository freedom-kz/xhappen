package boss

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	golog "log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gobwas/ws"
)

func WsServe(listener net.Listener, handler ConnHandler, logger log.Logger) error {
	logger.Log(log.LevelInfo, "ws.listener.addr", listener.Addr())
	var wg sync.WaitGroup

	http.HandleFunc("/im", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		defer wg.Done()
		handle(w, r, handler, logger)
	})

	server := &http.Server{
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,                                   //1M
		ErrorLog:       golog.New(os.Stderr, "", golog.LstdFlags), //打印底层错误日志至stderr
	}

	if err := server.Serve(listener); err != nil && strings.Contains(err.Error(), "use of closed network connection") {
		logger.Log(log.LevelInfo, "listener.close", listener.Addr())
	} else {
		logger.Log(log.LevelError, "err", err)
	}
	wg.Wait()
	return nil
}

func handle(w http.ResponseWriter, r *http.Request, handler ConnHandler, logger log.Logger) {
	upgrader := ws.HTTPUpgrader{
		Timeout: HandshakeTimeout,
	}
	conn, rwf, _, err := upgrader.Upgrade(r, w)
	if err != nil {
		logger.Log(log.LevelInfo, "msg", "upgrader ws fail", "err", err)
		return
	}
	//conn.SetReadLimit(MaxMessageSize)
	conn.SetReadDeadline(time.Now().Add(DEFAULT_READ_TIMEOUT))
	rwf = bufio.NewReadWriter(
		bufio.NewReaderSize(rwf, ReadBufferSize),
		bufio.NewWriterSize(rwf, WriteBufferSize),
	)
	wsConn := &WsConn{Conn: conn, RWB: rwf}
	handler.Handle(wsConn)
}

type WsConn struct {
	net.Conn
	Transform io.Reader
	RWB       *bufio.ReadWriter
}

func (wsConn *WsConn) Read(b []byte) (n int, err error) {
	if wsConn.Transform == nil {
		err := wsConn.SetReader()
		if err != nil {
			return 0, err
		}
	}

	n, err = wsConn.Transform.Read(b)
	if err == io.EOF {
		wsConn.Transform = nil
		return wsConn.Read(b)
	} else {
		return n, err
	}
}

func (wsConn *WsConn) SetReader() (err error) {
	frame, err := ws.ReadFrame(wsConn.RWB)
	if err != nil {
		return err
	}
	if frame.Header.OpCode != ws.OpBinary {
		return fmt.Errorf("websocket read opcode not binary:")
	}
	frame = ws.UnmaskFrameInPlace(frame)
	bts := frame.Payload
	wsConn.Transform = bytes.NewReader(bts)
	return nil
}

func (wsConn *WsConn) Write(bts []byte) (n int, err error) {
	frame := ws.NewBinaryFrame(bts)
	err = ws.WriteFrame(wsConn.RWB, frame)
	if err != nil {
		return 0, err
	}
	return len(bts), nil
}

func (ws *WsConn) Close() error {
	return ws.Conn.Close()
}

func (ws *WsConn) LocalAddr() net.Addr {
	return ws.Conn.LocalAddr()
}

func (ws *WsConn) RemoteAddr() net.Addr {
	return ws.Conn.RemoteAddr()
}

func (ws *WsConn) SetDeadline(t time.Time) error {
	err := ws.Conn.SetReadDeadline(t)
	if err != nil {
		return err
	}
	return ws.Conn.SetWriteDeadline(t)
}

func (ws *WsConn) SetReadDeadline(t time.Time) error {
	return ws.Conn.SetReadDeadline(t)
}

func (ws *WsConn) SetWriteDeadline(t time.Time) error {
	return ws.Conn.SetWriteDeadline(t)
}

func (ws *WsConn) Flush() error {
	return ws.RWB.Flush()
}
