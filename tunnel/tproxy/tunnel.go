// +build linux

package tproxy

import (
	"context"

	"github.com/frainzy1477/trojan-go/tunnel"
)

const Name = "TPROXY"

type Tunnel struct{}

func (t *Tunnel) Name() string {
	return Name
}

func (t *Tunnel) NewClient(ctx context.Context, client tunnel.Client) (tunnel.Client, error) {
	panic("not supported")
}

func (t *Tunnel) NewServer(ctx context.Context, server tunnel.Server) (tunnel.Server, error) {
	return NewServer(ctx, server)
}

func init() {
	tunnel.RegisterTunnel(Name, &Tunnel{})
}
