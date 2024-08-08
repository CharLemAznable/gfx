package gsse

import (
	"context"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gmutex"
)

const (
	NoEvent      = ""
	NoId         = ""
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

// SendMessage calls emit(noEvent, "data", noId)
func (c *Client) SendMessage(data string) {
	c.emit(NoEvent, data, NoId)
}

// SendMessageWithId calls emit(noEvent, "data", "id")
func (c *Client) SendMessageWithId(data, id string) {
	c.emit(NoEvent, data, id)
}

// SendEvent calls emit("event", "data", noId)
func (c *Client) SendEvent(event, data string) {
	c.emit(event, data, NoId)
}

// SendEventWithId calls emit("event", "data", "id")
func (c *Client) SendEventWithId(event, data, id string) {
	c.emit(event, data, id)
}

func (c *Client) emit(event, data, id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	select {
	case <-c.Context().Done():
	default:
		if event != NoEvent { // event: not required
			c.Response().Writeln("event:", event)
		}
		c.Response().Writeln("data:", data)
		if id != NoId { // id: not required
			c.Response().Writeln("id:", id)
		}
		c.Response().Writeln()
		c.Response().Flush()
	}
}

// SendComment send comment with prefix":"
func (c *Client) SendComment(comment string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.comment(comment)
}

func (c *Client) heartbeat() {
	c.mutex.TryLockFunc(func() {
		c.comment(EmptyComment)
	})
}

func (c *Client) comment(comment string) {
	select {
	case <-c.Context().Done():
	default:
		c.Response().Writeln(":", comment)
		c.Response().Writeln()
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
			go g.TryCatch(context.Background(), func(ctx context.Context) {
				onClose.(func(*Client))(client)
			}, func(ctx context.Context, exception error) {
				// ignore
			})
		}
	}()
	return client
}
