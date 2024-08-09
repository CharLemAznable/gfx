package gsse_test

import (
	"fmt"
	"github.com/CharLemAznable/gfx/net/gsse"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"testing"
	"time"
)

func Test_SendMessage(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/sse", gsse.Handle(func(client *gsse.Client) {
		client.SendMessage("send message1", "send message2")
		client.SendComment("send comment1", "send comment2")
		client.SendMessage()
		client.SendComment()
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := g.Client()
		client.SetPrefix(prefix)

		t.Assert(client.GetContent(gctx.New(), "/sse"),
			"data:send message1\ndata:send message2\n\n:send comment1\n:send comment2\n")
	})
}

func Test_SendMessageWithId(t *testing.T) {
	ch := make(chan bool, 1)
	s := g.Server(guid.S())
	s.BindHandler("/sse", gsse.Handle(func(client *gsse.Client) {
		client.OnClose(func(client *gsse.Client) {
			ch <- client.Terminated()
		})
		client.SendMessageWithId("1", "send message with id")
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := g.Client()
		client.SetPrefix(prefix)

		t.Assert(client.GetContent(gctx.New(), "/sse"),
			"id:1\ndata:send message with id\n\n")

		select {
		case value := <-ch:
			t.AssertEQ(value, true)
		}
	})
}

func Test_SendEvent(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/sse", gsse.Handle(func(client *gsse.Client) {
		client.KeepAlive()
		go func(client *gsse.Client) {
			<-time.After(time.Second)
			client.SendEvent("test", "send event")
			client.Close()
			client.SendEvent("test", "send event")
		}(client)
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := g.Client()
		client.SetPrefix(prefix)

		t.Assert(client.GetContent(gctx.New(), "/sse"),
			":\nevent:test\ndata:send event\n\n")
	})
}

func Test_SendEventWithId(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/sse", gsse.Handle(func(client *gsse.Client) {
		client.OnClose(func(client *gsse.Client) {
			panic("ignored")
		})
		client.KeepAlive()
		go func(client *gsse.Client) {
			<-time.After(time.Second)
			client.SendEventWithId("2", "test", "send event")
			client.Close()
			client.SendComment("send comment")
		}(client)
	}))
	s.SetDumpRouterMap(false)
	_ = s.Start()
	defer func() { _ = s.Shutdown() }()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := g.Client()
		client.SetPrefix(prefix)

		t.Assert(client.GetContent(gctx.New(), "/sse"),
			":\nid:2\nevent:test\ndata:send event\n\n")
	})
}
