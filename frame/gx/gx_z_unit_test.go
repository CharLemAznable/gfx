package gx_test

import (
	"github.com/CharLemAznable/gfx/frame/gx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"testing"
	"time"
)

func Test_Object(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gx.Client(), nil)
		t.AssertNE(gx.Config(), nil)
		t.AssertNE(gx.Cfg(), nil)
	})
}

func Test_Func(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
	})
	s.SetDumpRouterMap(false)
	gx.ConfigServer(s,
		gx.WithRandomAddr(),
		gx.WithDefaultHttpAddr(),
		gx.WithDefaultAddr(":8080"))
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(s.GetListenedPort(), 8080)
	})
}
