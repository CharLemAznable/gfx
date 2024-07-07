package ghttpx

import (
	"fmt"
	"github.com/CharLemAznable/gfx/os/gcmdx"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ConfigServer(server *ghttp.Server, configs ...func(server *ghttp.Server)) *ghttp.Server {
	for _, config := range configs {
		config(server)
	}
	return server
}

const (
	cmdEnvKeyForDefaultAddress = "gf.ghttp.server.address"
	cmdEnvKeyFormatForAddress  = "gf.ghttp.server.%s.address"
)

func WithDefaultAddr(address string) func(server *ghttp.Server) {
	return func(server *ghttp.Server) {
		var addrVar *gvar.Var
		serverName := server.GetName()
		if serverName == ghttp.DefaultServerName {
			addrVar = gcmdx.GetOptWithEnv(cmdEnvKeyForDefaultAddress, address)
		} else {
			key := fmt.Sprintf(cmdEnvKeyFormatForAddress, serverName)
			addrVar = gcmdx.GetOptWithEnv(key, address)
		}
		server.SetAddr(addrVar.String())
	}
}

func WithRandomAddr() func(server *ghttp.Server) {
	return WithDefaultAddr(":0")
}

func WithDefaultHttpAddr() func(server *ghttp.Server) {
	return WithDefaultAddr(":80")
}
