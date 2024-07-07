package gx

import (
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/CharLemAznable/gfx/os/gcfgx"
	"github.com/gogf/gf/v2/frame/g"
)

func Client() *gclientx.Client {
	return gclientx.New(g.Client())
}

func Config(name ...string) *gcfgx.Config {
	return gcfgx.New(g.Config(name...))
}

func Cfg(name ...string) *gcfgx.Config {
	return Config(name...)
}
