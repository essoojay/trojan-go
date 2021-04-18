package main

import (
	"flag"

	"github.com/frainzy1477/trojan-goo/option"
	
	_ "github.com/frainzy1477/trojan-goo/component"
	"github.com/frainzy1477/trojan-goo/log"
	
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
