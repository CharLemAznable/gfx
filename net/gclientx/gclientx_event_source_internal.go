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
	closedDone = make(chan struct{})
)

func init() {
	close(closedDone)
}

type internalEventSource struct {
	client *Client
	method string
	url    string
	data   []interface{}
	mutex  *gmutex.Mutex
	buffer chan *Event
	done   *gtype.Interface
	err    error
}

func newEventSource(client *Client, method string, url string, data ...interface{}) *internalEventSource {
	return &internalEventSource{
		client: client,
		method: method,
		url:    url,
		data:   data,
		mutex:  &gmutex.Mutex{},
		buffer: make(chan *Event, 1024),
		done:   gtype.NewInterface(),
	}
}

func (s *internalEventSource) Execute(listener ...EventListener) EventSource {
	if len(listener) > 0 && listener[0] != nil {
		go func(listener EventListener) {
			for event := range s.buffer {
				listener.OnEvent(event)
			}
			listener.OnClose(s.Err())
			s.finish()
		}(listener[0])
	} else {
		go func() {
			for range s.buffer {
				// drain the buffer
			}
			s.finish()
		}()
	}
	err := grpool.AddWithRecover(context.Background(), func(ctx context.Context) {
		response, err := s.client.Client.
			DoRequest(ctx, s.method, s.url, s.data...)
		if err != nil {
			s.close(err)
			return
		}
		defer s.client.deferCloseResponse(ctx, response)
		if response.StatusCode != http.StatusOK {
			s.close(gerror.New(string(response.ReadAll())))
			return
		}
		scanner := bufio.NewScanner(response.Body)
		for s.processNextEvent(scanner) {
		}
		s.close(scanner.Err())
	}, s.client.deferLogError)
	s.client.deferLogError(context.Background(), err)
	return s
}

func (s *internalEventSource) Done() <-chan struct{} {
	d := s.done.Val()
	if d != nil {
		return d.(chan struct{})
	}
	s.mutex.LockFunc(func() {
		d = s.done.Val()
		if d == nil {
			d = make(chan struct{})
			s.done.Set(d)
		}
	})
	return d.(chan struct{})
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
		close(s.buffer)
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
				s.buffer <- event
				return true
			}
		}
	}
	return false
}

func (s *internalEventSource) finish() {
	s.mutex.LockFunc(func() {
		d, _ := s.done.Val().(chan struct{})
		if d == nil {
			s.done.Set(closedDone)
		} else {
			close(d)
		}
	})
}
