package gsse

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"
)

// Handle wraps func(*gsse.Client) to func(*ghttp.Request), use by ghttp.Server.BindHandler().
func Handle(fn func(*Client)) func(*ghttp.Request) {
	return func(request *ghttp.Request) {
		client := newClient(request)
		if fn != nil {
			fn(client)
		}

		if !client.keepAlive {
			return
		}

		keepAliveCtx, keepAliveCancel :=
			context.WithCancel(context.Background())
		go func() {
			<-client.Context().Done()
			keepAliveCancel()
		}()
		for {
			select {
			case <-keepAliveCtx.Done():
				return
			case <-time.After(5 * time.Second):
				client.heartbeat()
			}
		}
	}
}
