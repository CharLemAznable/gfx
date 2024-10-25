package gclientx_test

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/CharLemAznable/gfx/net/gsse"
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"net/http"
	"sync"
	"testing"
	"time"
)

func Test_EventSource(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/sse", gsse.Handle(func(client *gsse.Client) {
		client.SendComment("send message")
		client.SendEventWithId("1", "message", " send message ", " send message ")
		client.SendComment("send message")
		client.SendEventWithId("1", "message", " send message ", " send message ")
		client.SendComment("send message")
		client.SendEventWithId("1", "message", " send message ", " send message ")
		client.SendComment("send message")
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()
	time.Sleep(100 * time.Millisecond)
	prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
	client := gclientx.New(g.Client()).Prefix(prefix)

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.GetEventSource(context.Background(), "/sse")
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			count := 0
			for {
				event, ok := <-eventSource.Event()
				count++
				t.AssertNil(eventSource.Err())
				if !ok {
					break
				}
				t.Assert(event.Event, "message")
				t.Assert(event.Data, "send message \nsend message ")
				t.Assert(event.Id, "1")
			}
			t.Assert(count, 4)
		}()
		wg.Wait()
	})

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.GetEventSource(context.Background(), "/sse")
		eventSource.Close()
		t.AssertNil(eventSource.Err())
	})
}

func Test_EventSource_Error(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/sse", gsse.Handle(func(client *gsse.Client) {
		client.SendEventWithId("1", "message", "send message")
		<-time.After(3 * time.Second)
		client.SendEventWithId("1", "message", "send message")
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()
	time.Sleep(100 * time.Millisecond)
	prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
	client := gclientx.New(g.Client())

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.GetEventSource(context.Background(), "")
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "request failed: Get \"http:\": http: no Host in request URL")
	})

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.Prefix(prefix).GetEventSource(context.Background(), "/notfound")
		eventSource.Close()
		httpErr, ok := eventSource.Err().(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), fmt.Sprintf("%d %s",
			http.StatusNotFound, http.StatusText(http.StatusNotFound)))
		t.Assert(httpErr.StatusCode(), http.StatusNotFound)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusNotFound))
	})

	gtest.C(t, func(t *gtest.T) {
		timeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		eventSource := client.Prefix(prefix).GetEventSource(timeout, "/sse")
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "context deadline exceeded")
	})
}

func Test_eventSource_Tmpl_Request(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/hello", gsse.Handle(func(client *gsse.Client) {
		client.SendComment("send message")
		client.SendEventWithId("1", "message", " send message ", " send message ")
		client.SendComment("send message")
		client.SendEventWithId("1", "message", " send message ", " send message ")
		client.SendComment("send message")
		client.SendEventWithId("1", "message", " send message ", " send message ")
		client.SendComment("send message")
	}))
	s.BindHandler("/error", gsse.Handle(func(client *gsse.Client) {
		client.Response().WriteStatusExit(http.StatusInternalServerError)
	}))
	s.BindHandler("/timeout", gsse.Handle(func(client *gsse.Client) {
		client.SendEventWithId("1", "message", "send message")
		<-time.After(3 * time.Second)
		client.SendEventWithId("1", "message", "send message")
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()
	time.Sleep(100 * time.Millisecond)
	client := gclientx.New()
	url := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
	view := gviewx.New().SetAdapter(gviewx.NewAdapterFile("testdata"))
	params := g.Map{"ListenedPort": s.GetListenedPort()}

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.RawContentEventSource(context.Background(), "GET "+url+"/hello HTTP/1.1\n\n")
		defer eventSource.Close()
		for event := range eventSource.Event() {
			t.Assert(event.Event, "message")
			t.Assert(event.Data, "send message \nsend message ")
			t.Assert(event.Id, "1")
		}
	})
	gtest.C(t, func(t *gtest.T) {
		eventSource := client.RawContentEventSource(context.Background(), "GET "+url+"/error HTTP/1.1\n\n")
		eventSource.Close()
		httpErr, ok := eventSource.Err().(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), fmt.Sprintf("%d %s",
			http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)))
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))
	})
	gtest.C(t, func(t *gtest.T) {
		eventSource := client.RawContentEventSource(context.Background(), "GET /hello HTTP/1.1\n\n")
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "request failed: Get \"/hello\": unsupported protocol scheme \"\"")
	})
	gtest.C(t, func(t *gtest.T) {
		timeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		eventSource := client.RawContentEventSource(timeout, "GET "+url+"/timeout HTTP/1.1\n\n")
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "context deadline exceeded")
	})

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.TmplEventSource(context.Background(), view, "hello", params)
		defer eventSource.Close()
		for event := range eventSource.Event() {
			t.Assert(event.Event, "message")
			t.Assert(event.Data, "send message \nsend message ")
			t.Assert(event.Id, "1")
		}
	})
	gtest.C(t, func(t *gtest.T) {
		eventSource := client.TmplEventSource(context.Background(), view, "error", params)
		eventSource.Close()
		httpErr, ok := eventSource.Err().(gclientx.HttpError)
		t.Assert(ok, true)
		t.Assert(httpErr.Error(), fmt.Sprintf("%d %s",
			http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)))
		t.Assert(httpErr.StatusCode(), http.StatusInternalServerError)
		t.Assert(httpErr.StatusText(), http.StatusText(http.StatusInternalServerError))
	})
	gtest.C(t, func(t *gtest.T) {
		eventSource := client.TmplEventSource(context.Background(), view, "fail", params)
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "request failed: Get \"/hello\": unsupported protocol scheme \"\"")
	})
	gtest.C(t, func(t *gtest.T) {
		timeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		eventSource := client.TmplEventSource(timeout, view, "timeout", params)
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "context deadline exceeded")
	})
}
