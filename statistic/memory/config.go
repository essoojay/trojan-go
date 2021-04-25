package memory

import (
	"github.com/frainzy1477/trojan-go/config"
)

type Config struct {
	Passwords []string `json:"password" yaml:"password"`
}

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return &Config{}
	})
}
