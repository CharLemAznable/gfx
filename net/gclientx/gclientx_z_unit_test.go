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

func Test_ClientX_Private(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		client := gclientx.New().
			Prefix("http://a.b.c").
			HeaderMap(map[string]string{"hKey": "hValue"}).
			CookieMap(map[string]string{"cKey": "cValue"}).
			BasicAuth("user", "pass")
		t.Assert(client.GetPrefix(), "http://a.b.c")
		t.Assert(client.GetHeader("hKey"), "hValue")
		t.Assert(client.GetCookie("cKey"), "cValue")
		user, pass := client.GetBasicAuth()
		t.Assert(user, "user")
		t.Assert(pass, "pass")

		// get map read-only copy
		client.GetHeaderMap()["hKey"] = "hValue2"
		client.GetCookieMap()["cKey"] = "cValue2"
		t.Assert(client.GetHeader("hKey"), "hValue")
		t.Assert(client.GetCookie("cKey"), "cValue")
	})
}

func Test_ClientX_Tmpl_Request(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("world")
		if hv := r.GetHeader("hKey"); hv != "" {
			r.Response.Writeln()
			r.Response.Write(hv)
		}
		if cv := r.Cookie.Get("cKey").String(); cv != "" {
			r.Response.Writeln()
			r.Response.Write(cv)
		}
	})
	s.BindHandler("/error", func(r *ghttp.Request) {
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	})
	s.BindHandler("/auth", func(r *ghttp.Request) {
		if r.BasicAuth("john", "123456") {
			r.Response.Write("ok")
		} else {
			r.Response.Write("fail")
		}
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

		content, err = client.Prefix(url).
			Header("hKey", "hValue").
			Header("Host", "127.0.0.1").
			Cookie("cKey", "cValue").
			Cookie("cKey1", "cValue1").
			RawContentRequestContent(ctx, "GET /hello HTTP/1.1\n\n")
		t.Assert(content, "world\nhValue\ncValue")
		t.AssertNil(err)

		_, err = client.Prefix(" ").DoRawContentRequest(ctx, "GET /hello HTTP/1.1\n\n")
		t.Assert(err.Error(), "prefix request failed: parse \"http:// /hello\": invalid character \" \" in host name")

		content, err = client.Prefix(url).BasicAuth("john", "123456").
			RawContentRequestContent(ctx, "GET /auth HTTP/1.1\n\n")
		t.Assert(content, "ok")
		t.AssertNil(err)
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

		content, err = client.Prefix(url).
			Header("hKey", "hValue").
			Header("Host", "127.0.0.1").
			Cookie("cKey", "cValue").
			Cookie("cKey1", "cValue1").
			TmplRequestContent(ctx, view, "fail")
		t.Assert(content, "world\nhValue\ncValue")
		t.AssertNil(err)

		_, err = client.Prefix(" ").DoTmplRequest(ctx, view, "fail")
		t.Assert(err.Error(), "prefix request failed: parse \"http:// /hello\": invalid character \" \" in host name")

		content, err = client.Prefix(url).BasicAuth("john", "123456").
			TmplRequestContent(ctx, view, "auth")
		t.Assert(content, "ok")
		t.AssertNil(err)
	})
}
