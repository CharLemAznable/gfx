package gclientx_test

import (
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
		eventSource := client.GetEventSource("/sse")
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
		eventSource := client.GetEventSource("/sse")
		eventSource.Close()
		t.AssertNil(eventSource.Err())
	})
}

func Test_EventSource_Error(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/sse", gsse.Handle(func(client *gsse.Client) {
		client.SendEventWithId("1", "message", "send message")
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()
	time.Sleep(100 * time.Millisecond)
	prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
	client := gclientx.New(g.Client())

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.GetEventSource("")
		eventSource.Close()
		t.AssertNE(eventSource.Err(), nil)
	})

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.Prefix(prefix).GetEventSource("/notfound")
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "Not Found")
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
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()
	time.Sleep(100 * time.Millisecond)
	client := gclientx.New()
	url := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
	view := gviewx.New().SetAdapter(gviewx.NewAdapterFile("testdata"))
	params := g.Map{"ListenedPort": s.GetListenedPort()}

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.RawContentEventSource("GET " + url + "/hello HTTP/1.1\n\n")
		defer eventSource.Close()
		for event := range eventSource.Event() {
			t.Assert(event.Event, "message")
			t.Assert(event.Data, "send message \nsend message ")
			t.Assert(event.Id, "1")
		}
	})
	gtest.C(t, func(t *gtest.T) {
		eventSource := client.RawContentEventSource("GET " + url + "/error HTTP/1.1\n\n")
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "Internal Server Error")
	})
	gtest.C(t, func(t *gtest.T) {
		eventSource := client.RawContentEventSource("GET /hello HTTP/1.1\n\n")
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "request failed: Get \"/hello\": unsupported protocol scheme \"\"")
	})

	gtest.C(t, func(t *gtest.T) {
		eventSource := client.TmplEventSource(view, "hello", params)
		defer eventSource.Close()
		for event := range eventSource.Event() {
			t.Assert(event.Event, "message")
			t.Assert(event.Data, "send message \nsend message ")
			t.Assert(event.Id, "1")
		}
	})
	gtest.C(t, func(t *gtest.T) {
		eventSource := client.TmplEventSource(view, "error", params)
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "Internal Server Error")
	})
	gtest.C(t, func(t *gtest.T) {
		eventSource := client.TmplEventSource(view, "fail", params)
		eventSource.Close()
		t.Assert(eventSource.Err().Error(), "request failed: Get \"/hello\": unsupported protocol scheme \"\"")
	})
}
