package gx

import (
	"github.com/CharLemAznable/gfx/net/ghttpx"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ConfigServer(server *ghttp.Server, configs ...func(server *ghttp.Server)) *ghttp.Server {
	return ghttpx.ConfigServer(server, configs...)
}

func WithDefaultAddr(address string) func(server *ghttp.Server) {
	return ghttpx.WithDefaultAddr(address)
}

func WithRandomAddr() func(server *ghttp.Server) {
	return ghttpx.WithRandomAddr()
}

func WithDefaultHttpAddr() func(server *ghttp.Server) {
	return ghttpx.WithDefaultHttpAddr()
}
