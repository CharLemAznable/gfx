package gx

import (
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/CharLemAznable/gfx/net/ghttpx"
	"github.com/CharLemAznable/gfx/os/gcfgx"
	"github.com/gogf/gf/v2/frame/g"
)

func Client() *gclientx.Client {
	return gclientx.New(g.Client())
}

func Server(name ...interface{}) *ghttpx.Server {
	return ghttpx.GetServer(name...)
}

func Config(name ...string) *gcfgx.Config {
	return gcfgx.Instance(name...)
}

func Cfg(name ...string) *gcfgx.Config {
	return Config(name...)
}
