package gclientx_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"net/http"
	"testing"
	"time"
)

var (
	ctx = context.TODO()
)

func Test_ClientX(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
	})
	s.BindHandler("/error", func(r *ghttp.Request) {
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	})
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := gclientx.New(g.Client())

		bytes, err := client.GetBytes(ctx, "")
		t.AssertNil(bytes)
		t.AssertNE(err, nil)
		bytes, err = client.PostBytes(ctx, "")
		t.AssertNil(bytes)
		t.AssertNE(err, nil)
		content, err := client.GetContent(ctx, "")
		t.Assert(content, "")
		t.AssertNE(err, nil)
		content, err = client.PostContent(ctx, "")
		t.Assert(content, "")
		t.AssertNE(err, nil)
		v, err := client.GetVar(ctx, "")
		t.AssertNil(v.Val())
		t.AssertNE(err, nil)
		v, err = client.PostVar(ctx, "")
		t.AssertNil(v.Val())
		t.AssertNE(err, nil)

		url := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client = client.Prefix(url)
		bytes, err = client.GetBytes(ctx, "/hello")
		t.Assert(bytes, []byte("world"))
		t.AssertNil(err)
		bytes, err = client.PostBytes(ctx, "/hello")
		t.Assert(bytes, []byte("world"))
		t.AssertNil(err)
		content, err = client.GetContent(ctx, "/hello")
		t.Assert(content, "world")
		t.AssertNil(err)
		content, err = client.PostContent(ctx, "/hello")
		t.Assert(content, "world")
		t.AssertNil(err)
		v, err = client.GetVar(ctx, "/hello")
		t.Assert(v.Val(), "world")
		t.AssertNil(err)
		v, err = client.PostVar(ctx, "/hello")
		t.Assert(v.Val(), "world")
		t.AssertNil(err)

		errStr := fmt.Sprintf("%d: %s",
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		bytes, err = client.GetBytes(ctx, "/error")
		t.AssertNil(bytes)
		t.Assert(err.Error(), errStr)
		bytes, err = client.PostBytes(ctx, "/error")
		t.AssertNil(bytes)
		t.Assert(err.Error(), errStr)
		content, err = client.GetContent(ctx, "/error")
		t.Assert(content, "")
		t.Assert(err.Error(), errStr)
		content, err = client.PostContent(ctx, "/error")
		t.Assert(content, "")
		t.Assert(err.Error(), errStr)
		v, err = client.GetVar(ctx, "/error")
		t.AssertNil(v.Val())
		t.Assert(err.Error(), errStr)
		v, err = client.PostVar(ctx, "/error")
		t.AssertNil(v.Val())
		t.Assert(err.Error(), errStr)
	})
}

func Test_ClientX_SetIntLog(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gclientx.New().SetIntLog(g.Log()), nil)
		t.AssertNE(gclientx.New().SetIntLog(nil), nil)
	})
}
