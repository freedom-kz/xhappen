package protocol

import (
	"bufio"
	"net"
	"time"
)

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
