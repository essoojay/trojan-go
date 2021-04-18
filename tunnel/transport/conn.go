package transport

import (
	"net"

	"github.com/frainzy1477/trojan-goo/tunnel"
)

type Conn struct {
	net.Conn
}

func (c *Conn) Metadata() *tunnel.Metadata {
	return nil
}
