package gsse

import (
	"context"
	"github.com/CharLemAznable/gfx/frame/gx"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gmutex"
)

const (
	NoId         = ""
	NoEvent      = ""
	EmptyComment = ""
)

// Client wraps the SSE(Server-Sent Event) ghttp.Request and provides SSE APIs
type Client struct {
	request   *ghttp.Request
	cancel    context.CancelFunc
	onClose   *gtype.Interface
	keepAlive bool
	mutex     *gmutex.Mutex
}

// Request return ghttp.Request.
func (c *Client) Request() *ghttp.Request {
	return c.request
}

// Response is alias for ghttp.Request.Response.
func (c *Client) Response() *ghttp.Response {
	return c.Request().Response
}

// Context is alias for ghttp.Request.Context().
func (c *Client) Context() context.Context {
	return c.Request().Context()
}

// SendMessage calls emit(NoId, NoEvent, data...)
func (c *Client) SendMessage(data ...string) {
	c.emit(NoId, NoEvent, data...)
}

// SendMessageWithId calls emit(id, NoEvent, data...)
func (c *Client) SendMessageWithId(id string, data ...string) {
	c.emit(id, NoEvent, data...)
}

// SendEvent calls emit(NoId, event, data...)
func (c *Client) SendEvent(event string, data ...string) {
	c.emit(NoId, event, data...)
}

// SendEventWithId calls emit(id, event, data...)
func (c *Client) SendEventWithId(id, event string, data ...string) {
	c.emit(id, event, data...)
}

func (c *Client) emit(id, event string, data ...string) {
	if len(data) == 0 { // data: required
		return
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	select {
	case <-c.Context().Done():
	default:
		if id != NoId { // id: not required
			c.Response().Writeln("id:", id)
		}
		if event != NoEvent { // event: not required
			c.Response().Writeln("event:", event)
		}
		for _, dt := range data {
			c.Response().Writeln("data:", dt)
		}
		c.Response().Writeln()
		c.Response().Flush()
	}
}

// SendComment send comment with prefix":"
func (c *Client) SendComment(comment ...string) {
	if len(comment) == 0 {
		return
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.comment(comment...)
}

func (c *Client) heartbeat() {
	c.mutex.TryLockFunc(func() {
		c.comment(EmptyComment)
	})
}

func (c *Client) comment(comment ...string) {
	select {
	case <-c.Context().Done():
	default:
		for _, cm := range comment {
			c.Response().Writeln(":", cm)
		}
		c.Response().Flush()
	}
}

// Close closes the connection
func (c *Client) Close() {
	c.cancel()
}

// Terminated returns true if the connection has been closed
func (c *Client) Terminated() bool {
	return c.Context().Err() != nil
}

// OnClose callback which runs when a client closes its connection
func (c *Client) OnClose(fn func(*Client)) {
	c.onClose.Set(fn)
}

// KeepAlive keeps the connection alive, if you need to use the client outside the handler
func (c *Client) KeepAlive() {
	c.keepAlive = true
}

func newClient(request *ghttp.Request) *Client {
	ctx, cancel := context.WithCancel(request.Context())
	request.SetCtx(ctx)
	request.Response.Header().Set("Content-Type", "text/event-stream")
	request.Response.Header().Set("Cache-Control", "no-cache")
	request.Response.Header().Set("Connection", "keep-alive")
	client := &Client{
		request:   request,
		cancel:    cancel,
		onClose:   gtype.NewInterface(),
		keepAlive: false,
		mutex:     &gmutex.Mutex{},
	}
	go func() {
		<-client.Context().Done()
		if onClose := client.onClose.Val(); onClose != nil {
			gx.GoAnywayX(func() {
				onClose.(func(*Client))(client)
			})
		}
	}()
	return client
}
