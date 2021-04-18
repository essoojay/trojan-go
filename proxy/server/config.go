package server

import (
	"github.com/frainzy1477/trojan-goo/config"
	"github.com/frainzy1477/trojan-goo/proxy/client"
)

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return new(client.Config)
	})
}
