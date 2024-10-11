package gclientx_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/CharLemAznable/gfx/os/gviewx"
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

		errStr := fmt.Sprintf("%d %s",
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		bytes, err = client.GetBytes(ctx, "/error")
		t.AssertNil(bytes)
		httpErr, ok := err.(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), errStr)
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))
		bytes, err = client.PostBytes(ctx, "/error")
		t.AssertNil(bytes)
		httpErr, ok = err.(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), errStr)
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))
		content, err = client.GetContent(ctx, "/error")
		t.Assert(content, "")
		httpErr, ok = err.(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), errStr)
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))
		content, err = client.PostContent(ctx, "/error")
		t.Assert(content, "")
		httpErr, ok = err.(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), errStr)
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))
		v, err = client.GetVar(ctx, "/error")
		t.AssertNil(v.Val())
		httpErr, ok = err.(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), errStr)
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))
		v, err = client.PostVar(ctx, "/error")
		t.AssertNil(v.Val())
		httpErr, ok = err.(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), errStr)
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))
	})
}

func Test_ClientX_SetIntLog(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gclientx.New().SetIntLog(g.Log()), nil)
		t.AssertNE(gclientx.New().SetIntLog(nil), nil)
	})
}

func Test_ClientX_Tmpl_Request(t *testing.T) {
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
	client := gclientx.New()
	url := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
	view := gviewx.New().SetAdapter(gviewx.NewAdapterFile("testdata"))
	params := g.Map{"ListenedPort": s.GetListenedPort()}

	gtest.C(t, func(t *gtest.T) {
		content, err := client.RawContentRequestContent(ctx, "GET "+url+"/hello HTTP/1.1\n\n")
		t.Assert(content, "world")
		t.AssertNil(err)
		val, err := client.RawContentRequestVar(ctx, "GET "+url+"/hello HTTP/1.1\n\n")
		t.Assert(val.String(), "world")
		t.AssertNil(err)

		bytes, err := client.RawContentRequestBytes(ctx, "GET "+url+"/error HTTP/1.1\n\n")
		t.AssertNil(bytes)
		httpErr, ok := err.(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))

		_, err = client.RawContentRequestVar(ctx, "GET "+url+"/illegal HTTP/1.1")
		t.Assert(err.Error(), "read request failed: unexpected EOF")

		_, err = client.DoRawContentRequest(ctx, "GET /hello HTTP/1.1\n\n")
		t.Assert(err.Error(), "request failed: Get \"/hello\": unsupported protocol scheme \"\"")
	})

	gtest.C(t, func(t *gtest.T) {
		content, err := client.TmplRequestContent(ctx, view, "hello", params)
		t.Assert(content, "world")
		t.AssertNil(err)
		val, err := client.TmplRequestVar(ctx, view, "hello", params)
		t.Assert(val.String(), "world")
		t.AssertNil(err)

		bytes, err := client.TmplRequestBytes(ctx, view, "error", params)
		t.AssertNil(bytes)
		httpErr, ok := err.(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))

		_, err = client.TmplRequestContent(ctx, view, "notfound", params)
		t.Assert(err.Error(), "parse tmpl failed: template file \"notfound\" not found")
		_, err = client.TmplRequestVar(ctx, view, "illegal", params)
		t.Assert(err.Error(), "read request failed: unexpected EOF")

		_, err = client.DoTmplRequest(ctx, view, "fail")
		t.Assert(err.Error(), "request failed: Get \"/hello\": unsupported protocol scheme \"\"")
	})
}
