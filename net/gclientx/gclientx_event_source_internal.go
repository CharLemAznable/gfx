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

type internalEventSource struct {
	client  *Client
	method  string
	url     string
	data    []interface{}
	mutex   *gmutex.Mutex
	eventCh *gtype.Interface // Event channel.
	eventLn *gtype.Interface // Event listener.
	err     error
}

func newEventSource(client *Client, method string, url string, data ...interface{}) *internalEventSource {
	return &internalEventSource{
		client:  client,
		method:  method,
		url:     url,
		data:    data,
		mutex:   &gmutex.Mutex{},
		eventCh: gtype.NewInterface(),
		eventLn: gtype.NewInterface(),
	}
}

func (s *internalEventSource) Execute(ctx context.Context, listener ...EventListener) (EventSource, error) {
	if len(listener) > 0 && listener[0] != nil {
		s.mutex.LockFunc(func() {
			s.eventLn.Set(listener[0])
		})
	}
	return s, grpool.AddWithRecover(ctx, func(ctx context.Context) {
		response, err := s.client.Client.
			DoRequest(ctx, s.method, s.url, s.data...)
		if err != nil {
			s.close(err)
			return
		}
		defer s.client.deferCloseResponse(ctx, response)
		if response.StatusCode != http.StatusOK {
			s.close(gerror.New(response.Status))
			return
		}
		scanner := bufio.NewScanner(response.Body)
		defer s.close(scanner.Err())
		for s.processNextEvent(ctx, scanner) {
		}
	}, s.client.deferLogError)
}

func (s *internalEventSource) Event() <-chan *Event {
	ch := s.eventCh.Val()
	if ch != nil {
		return ch.(chan *Event)
	}
	s.mutex.LockFunc(func() {
		ch = s.eventCh.Val()
		if ch == nil {
			ch = make(chan *Event)
			s.eventCh.Set(ch)
		}
	})
	return ch.(chan *Event)
}

func (s *internalEventSource) Err() (err error) {
	s.mutex.LockFunc(func() {
		err = s.err
	})
	return
}

func (s *internalEventSource) close(err error) {
	s.mutex.LockFunc(func() {
		s.err = err
		if ch := s.eventCh.Val(); ch != nil {
			close(ch.(chan *Event))
		} else {
			s.eventCh.Set(closedEvent)
		}
		if ln := s.eventLn.Val(); ln != nil {
			go ln.(EventListener).OnEvent(nil, err)
		}
	})
}

func (s *internalEventSource) processNextEvent(ctx context.Context, scanner *bufio.Scanner) bool {
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
				s.completeEvent(ctx, event)
				return true
			}
		}
	}
	return false
}

func (s *internalEventSource) completeEvent(ctx context.Context, event *Event) {
	if ch := s.eventCh.Val(); ch != nil {
		_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
			ch.(chan *Event) <- event
		}, s.client.deferLogError)
	}
	if ln := s.eventLn.Val(); ln != nil {
		_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
			ln.(EventListener).OnEvent(event, nil)
		}, s.client.deferLogError)
	}
}
