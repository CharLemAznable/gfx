package gclientx

import (
	"bufio"
	"context"
	"github.com/gogf/gf/v2/container/gqueue"
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
	queueCh *gqueue.Queue
	eventCh *gtype.Interface // Event channel.
	queueLn *gqueue.Queue
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
		queueCh: gqueue.New(),
		eventCh: gtype.NewInterface(),
		queueLn: gqueue.New(),
		eventLn: gtype.NewInterface(),
	}
}

func (s *internalEventSource) Execute(listener ...EventListener) (EventSource, error) {
	if len(listener) > 0 && listener[0] != nil {
		s.mutex.LockFunc(func() {
			s.eventLn.Set(listener[0])
			go func(listener EventListener) {
				for event := range s.queueCh.C {
					listener.OnEvent(event.(*Event), nil)
				}
				listener.OnEvent(nil, s.Err())
			}(listener[0])
		})
	}
	return s, grpool.AddWithRecover(context.Background(), func(ctx context.Context) {
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
		for s.processNextEvent(scanner) {
		}
		s.close(scanner.Err())
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
			go func(ch chan *Event) {
				for event := range s.queueCh.C {
					ch <- event.(*Event)
				}
				close(ch)
			}(ch.(chan *Event))
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

func (s *internalEventSource) Close() {
	for range s.Event() {
		// drain the event channel
	}
}

func (s *internalEventSource) close(err error) {
	s.mutex.LockFunc(func() {
		s.err = err
		s.queueCh.Close()
		if ch := s.eventCh.Val(); ch == nil {
			s.eventCh.Set(closedEvent)
		}
		s.queueLn.Close()
	})
}

func (s *internalEventSource) processNextEvent(scanner *bufio.Scanner) bool {
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
				s.queueCh.Push(event)
				s.queueLn.Push(event)
				return true
			}
		}
	}
	return false
}
