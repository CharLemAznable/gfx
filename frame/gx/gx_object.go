package gx

import (
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/CharLemAznable/gfx/net/ghttpx"
	"github.com/CharLemAznable/gfx/os/gcfgx"
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/net/gclient"
)

func Client(client ...*gclient.Client) *gclientx.Client {
	return gclientx.New(client...)
}

func Server(name ...interface{}) *ghttpx.Server {
	return ghttpx.GetServer(name...)
}

func ViewX(name ...string) *gviewx.View {
	return gviewx.Instance(name...)
}

func Config(name ...string) *gcfgx.Config {
	return gcfgx.Instance(name...)
}

func Cfg(name ...string) *gcfgx.Config {
	return Config(name...)
}
