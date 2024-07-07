package ghttpx_test

import (
	"github.com/CharLemAznable/gfx/net/ghttpx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
	"time"
)

func Test_ConfigServer_Default(t *testing.T) {
	s := g.Server()
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
	})
	s.SetDumpRouterMap(false)
	ghttpx.ConfigServer(s, ghttpx.WithDefaultAddr(":8080"))
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(s.GetListenedPort(), 8080)
	})
}

func Test_ConfigServer_Opt(t *testing.T) {
	_ = genv.Set("GF_GHTTP_SERVER_OPT_ADDRESS", ":8081")
	defer func() { _ = genv.Remove("GF_GHTTP_SERVER_OPT_ADDRESS") }()
	s := g.Server("opt")
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
	})
	s.SetDumpRouterMap(false)
	ghttpx.ConfigServer(s, ghttpx.WithRandomAddr())
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(s.GetListenedPort(), 8081)
	})
}

func Test_ConfigServer_Env(t *testing.T) {
	gcmd.Init([]string{"--gf.ghttp.server.env.address=:8082"}...)
	s := g.Server("env")
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
	})
	s.SetDumpRouterMap(false)
	ghttpx.ConfigServer(s, ghttpx.WithDefaultHttpAddr())
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(s.GetListenedPort(), 8082)
	})
}
