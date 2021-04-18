// +build custom full

package build

import (
	_ "github.com/frainzy1477/trojan-goo/proxy/custom"
	_ "github.com/frainzy1477/trojan-goo/tunnel/adapter"
	_ "github.com/frainzy1477/trojan-goo/tunnel/dokodemo"
	_ "github.com/frainzy1477/trojan-goo/tunnel/freedom"
	_ "github.com/frainzy1477/trojan-goo/tunnel/http"
	_ "github.com/frainzy1477/trojan-goo/tunnel/mux"
	_ "github.com/frainzy1477/trojan-goo/tunnel/router"
	_ "github.com/frainzy1477/trojan-goo/tunnel/shadowsocks"
	_ "github.com/frainzy1477/trojan-goo/tunnel/simplesocks"
	_ "github.com/frainzy1477/trojan-goo/tunnel/socks"
	_ "github.com/frainzy1477/trojan-goo/tunnel/tls"
	_ "github.com/frainzy1477/trojan-goo/tunnel/tproxy"
	_ "github.com/frainzy1477/trojan-goo/tunnel/transport"
	_ "github.com/frainzy1477/trojan-goo/tunnel/trojan"
	_ "github.com/frainzy1477/trojan-goo/tunnel/websocket"
)
