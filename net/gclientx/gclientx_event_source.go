package gclientx

import (
	"bufio"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

type Event struct {
	Id    string
	Event string
	Data  string
}

type EventSource interface {
	Event() <-chan *Event
	Err() error
	Close()
}

func (c *Client) EventSource(ctx context.Context, method string, url string, data ...interface{}) EventSource {
	s := newEventSource()
	g.Go(ctx, func(ctx context.Context) {
		response, err := c.Client.DoRequest(ctx, method, url, data...)
		if err != nil {
			s.close(err)
			return
		}
		defer c.deferCloseResponse(ctx, response)
		if statusCode := response.StatusCode; statusCode != http.StatusOK {
			s.close(NewHttpError(statusCode, string(response.ReadAll())))
			return
		}
		scanner := bufio.NewScanner(response.Body)
		defer func() { s.close(scanner.Err()) }()
		for s.processNextEvent(scanner) {
		}
	}, c.deferLogError)
	return s
}
