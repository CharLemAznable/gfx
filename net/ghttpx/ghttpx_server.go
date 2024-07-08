package ghttpx

import (
	"fmt"
	"github.com/CharLemAznable/gfx/os/gcmdx"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var serverMapping = gmap.NewStrAnyMap(true)

func GetServer(name ...interface{}) *Server {
	serverName := ghttp.DefaultServerName
	if len(name) > 0 && name[0] != "" {
		serverName = gconv.String(name[0])
	}
	return serverMapping.GetOrSetFuncLock(serverName, func() interface{} {
		return &Server{Server: g.Server(serverName)}
	}).(*Server)
}

type Server struct {
	*ghttp.Server
}

const (
	cmdEnvKeyForDefaultAddress = "gf.ghttp.server.address"
	cmdEnvKeyFormatForAddress  = "gf.ghttp.server.%s.address"
)

func (server *Server) SetDefaultAddr(address string) *Server {
	var addrVar *gvar.Var
	serverName := server.Server.GetName()
	if serverName == ghttp.DefaultServerName {
		addrVar = gcmdx.GetOptWithEnv(cmdEnvKeyForDefaultAddress, address)
	} else {
		key := fmt.Sprintf(cmdEnvKeyFormatForAddress, serverName)
		addrVar = gcmdx.GetOptWithEnv(key, address)
	}
	server.Server.SetAddr(addrVar.String())
	return server
}

func (server *Server) SetRandomAddr() *Server {
	return server.SetDefaultAddr(":0")
}

func (server *Server) SetDefaultHttpAddr() *Server {
	return server.SetDefaultAddr(":80")
}
