package gclientx_test

import (
	"fmt"
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/CharLemAznable/gfx/net/gsse"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
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
