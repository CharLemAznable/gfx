package ghttpx_test

import (
	"github.com/CharLemAznable/gfx/net/ghttpx"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
	"time"
)

func Test_Server_SetDefaultAddr(t *testing.T) {
	_ = genv.Set("GF_GHTTP_SERVER_ADDRESS", "")
	defer func() { _ = genv.Remove("GF_GHTTP_SERVER_ADDRESS") }()
	gcmd.Init([]string{"--gf.ghttp.server.address="}...)
	s := ghttpx.GetServer()
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
	})
	s.SetDumpRouterMap(false)
	s.SetDefaultAddr(":8080")
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(s.GetListenedPort(), 8080)
	})
}

func Test_Server_Env_SetRandomAddr(t *testing.T) {
	_ = genv.Set("GF_GHTTP_SERVER_ENV_ADDRESS", ":8081")
	defer func() { _ = genv.Remove("GF_GHTTP_SERVER_ENV_ADDRESS") }()
	s := ghttpx.GetServer("env")
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
	})
	s.SetDumpRouterMap(false)
	s.SetRandomAddr() // :0
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(s.GetListenedPort(), 8081)
	})
}

func Test_Server_Cmd_SetDefaultHttpAddr(t *testing.T) {
	gcmd.Init([]string{"--gf.ghttp.server.cmd.address=:8082"}...)
	s := ghttpx.GetServer("cmd")
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
	})
	s.SetDumpRouterMap(false)
	s.SetDefaultHttpAddr() // :80
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(s.GetListenedPort(), 8082)
	})
}
