package gclientx

import (
	"bufio"
	"context"
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/gogf/gf/v2/os/gview"
	"net/http"
)

func (c *Client) TmplEventSource(view *gviewx.View,
	key string, params ...gview.Params) EventSource {
	s := &internalEventSource{
		mutex:  &gmutex.Mutex{},
		buffer: make(chan *Event, 1024),
	}
	g.Go(context.Background(), func(ctx context.Context) {
		response, err := c.DoTmplRequest(ctx, view, key, params...)
		if err != nil {
			s.close(err)
			return
		}
		defer c.deferCloseRawResponse(ctx, response)
		if response.StatusCode != http.StatusOK {
			s.close(gerror.New(string(c.readAll(ctx, response))))
			return
		}
		scanner := bufio.NewScanner(response.Body)
		defer s.close(scanner.Err())
		for s.processNextEvent(scanner) {
		}
	}, c.deferLogError)
	return s
}
