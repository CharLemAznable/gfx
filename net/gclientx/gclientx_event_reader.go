package gclientx

import (
	"bufio"
	"context"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/gogf/gf/v2/os/grpool"
	"net/http"
	"strings"
)

var (
	closedEvent = make(chan *Event)
)

func init() {
	close(closedEvent)
}

type eventReader struct {
	client *Client
	method string
	url    string
	data   []interface{}
	mutex  *gmutex.Mutex
	event  *gtype.Interface
	err    error
}

func newEventReader(client *Client, method string, url string, data ...interface{}) *eventReader {
	return &eventReader{
		client: client,
		method: method,
		url:    url,
		data:   data,
		mutex:  &gmutex.Mutex{},
		event:  gtype.NewInterface(),
	}
}

func (r *eventReader) Execute(ctx context.Context) error {
	return grpool.Add(ctx, func(ctx context.Context) {
		response, err := r.client.Client.Header(map[string]string{
			"Accept":        "text/event-stream",
			"Cache-Control": "no-cache",
			"Connection":    "keep-alive",
		}).DoRequest(ctx, r.method, r.url, r.data...)
		if err != nil {
			r.close(err)
			return
		}
		defer r.client.deferCloseResponse(ctx, response)
		if response.StatusCode != http.StatusOK {
			r.close(gerror.New(response.Status))
			return
		}
		scanner := bufio.NewScanner(response.Body)
		for r.processNextEvent(scanner) {
		}
	})
}

func (r *eventReader) Event() <-chan *Event {
	ch := r.event.Val()
	if ch != nil {
		return ch.(chan *Event)
	}
	r.mutex.LockFunc(func() {
		ch = r.event.Val()
		if ch == nil {
			ch = make(chan *Event)
			r.event.Set(ch)
		}
	})
	return ch.(chan *Event)
}

func (r *eventReader) Err() (err error) {
	r.mutex.LockFunc(func() {
		err = r.err
	})
	return
}

func (r *eventReader) close(err error) {
	r.mutex.LockFunc(func() {
		r.err = err
		ch := r.event.Val()
		if ch == nil {
			r.event.Set(closedEvent)
		} else {
			close(ch.(chan *Event))
		}
	})
}

func (r *eventReader) processNextEvent(scanner *bufio.Scanner) bool {
	event := &Event{}
	foundEvent := false
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "id:"):
			event.Id = strings.TrimSpace(strings.TrimPrefix(line, "id:"))
		case strings.HasPrefix(line, "event:"):
			event.Event = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
		case strings.HasPrefix(line, "data:"):
			event.Data = strings.TrimSpace(strings.TrimPrefix(line, "data:"))
			foundEvent = true
		default:
			if strings.TrimSpace(line) == "" && foundEvent {
				r.completeEvent(event)
				return true
			}
		}
	}
	r.close(scanner.Err())
	return false
}

func (r *eventReader) completeEvent(event *Event) {
	if ch := r.event.Val(); ch != nil {
		ch.(chan *Event) <- event
	}
}
