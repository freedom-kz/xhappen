package protocol

import "net"

type client struct {
	net.Conn
}
