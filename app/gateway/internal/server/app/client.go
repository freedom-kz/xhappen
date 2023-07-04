package app

import "net"

type client struct {
	net.Conn
}
