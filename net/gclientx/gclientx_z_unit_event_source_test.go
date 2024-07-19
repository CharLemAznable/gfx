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
		client.SendComment("send message")
		client.SendEventWithId("message", "send message", "1")
		client.SendComment("send message")
		client.SendComment("send message")
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := gclientx.New(g.Client()).Prefix(prefix)

		eventSource := client.GetEventSource("/sse")
		t.AssertNE(eventSource.Event(), nil)
		err := eventSource.Execute(ctx)
		t.AssertNil(err)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			count := 0
			select {
			case event, ok := <-eventSource.Event():
				if !ok {
					t.AssertNil(eventSource.Err())
					break
				}
				t.Assert(event.Event, "message")
				t.Assert(event.Data, "send message")
				t.Assert(event.Id, "1")
				count++
			}
			t.Assert(count, 1)
		}()
		wg.Wait()
	})
}

func Test_EventSource_Error(t *testing.T) {
	ch := make(chan bool, 1)
	s := g.Server(guid.S())
	s.BindHandler("/sse", gsse.Handle(func(client *gsse.Client) {
		client.OnClose(func(client *gsse.Client) {
			time.Sleep(time.Second)
			ch <- client.Terminated()
		})
		client.SendEventWithId("message", "send message", "1")
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := gclientx.New(g.Client())

		eventSource := client.GetEventSource("")
		t.AssertNE(eventSource.Event(), nil)
		err := eventSource.Execute(ctx)
		t.AssertNil(err)

		select {
		case event, ok := <-eventSource.Event():
			t.AssertNil(event)
			t.Assert(ok, false)
			t.AssertNE(eventSource.Err(), nil)
		}

		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client = client.Prefix(prefix)

		eventSource = client.GetEventSource("/notfound")
		t.AssertNE(eventSource.Event(), nil)
		err = eventSource.Execute(ctx)
		t.AssertNil(err)

		select {
		case event, ok := <-eventSource.Event():
			t.AssertNil(event)
			t.Assert(ok, false)
			t.Assert(eventSource.Err().Error(), "404 Not Found")
		}

		eventSource = client.GetEventSource("/sse")
		t.AssertNE(eventSource, nil)
		err = eventSource.Execute(ctx)
		t.AssertNil(err)

		select {
		case value := <-ch:
			t.Assert(value, true)
		}

		select {
		case event, ok := <-eventSource.Event():
			t.AssertNil(event)
			t.Assert(ok, false)
			t.AssertNil(eventSource.Err())
		}
	})
}
