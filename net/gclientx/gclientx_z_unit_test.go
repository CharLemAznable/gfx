package gclientx_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
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
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := gclientx.New(g.Client())
		client.SetErrorFn(func(ctx context.Context, format string, v ...interface{}) {
			t.AssertNE(v[0], nil)
		})

		bytes, err := client.GetBytesErr(ctx, "")
		t.AssertNil(bytes)
		t.AssertNE(err, nil)
		bytes, err = client.PostBytesErr(ctx, "")
		t.AssertNil(bytes)
		t.AssertNE(err, nil)
		content, err := client.GetContentErr(ctx, "")
		t.Assert(content, "")
		t.AssertNE(err, nil)
		content, err = client.PostContentErr(ctx, "")
		t.Assert(content, "")
		t.AssertNE(err, nil)

		url := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client = client.Prefix(url)
		bytes, err = client.GetBytesErr(ctx, "/hello")
		t.Assert(bytes, []byte("world"))
		t.AssertNil(err)
		bytes, err = client.PostBytesErr(ctx, "/hello")
		t.Assert(bytes, []byte("world"))
		t.AssertNil(err)
		content, err = client.GetContentErr(ctx, "/hello")
		t.Assert(content, "world")
		t.AssertNil(err)
		content, err = client.PostContentErr(ctx, "/hello")
		t.Assert(content, "world")
		t.AssertNil(err)
	})
}

func Test_ClientX_SetErrorLogger(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gclientx.New().SetErrorLogger(g.Log()), nil)
		t.AssertNE(gclientx.New().SetErrorLogger(nil), nil)
	})
}
