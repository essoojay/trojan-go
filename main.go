package main

import (
	"flag"

	"github.com/frainzy1477/trojan-go/option"
	
	_ "github.com/frainzy1477/trojan-go/component"
	"github.com/frainzy1477/trojan-go/log"
	
)

func main() {
	flag.Parse()
	for {
		h, err := option.PopOptionHandler()
		if err != nil {
			log.Fatal("invalid options")
		}
		err = h.Handle()
		if err == nil {
			break
		}
	}
}
