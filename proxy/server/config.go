package server

import (
	"github.com/frainzy1477/trojan-go/config"
	"github.com/frainzy1477/trojan-go/proxy/client"
)

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return new(client.Config)
	})
}
