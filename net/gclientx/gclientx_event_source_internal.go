package gclientx

import (
	"bufio"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gmutex"
	"net/http"
	"strings"
)

type internalEventSource struct {
	client *Client
	method string
	url    string
	data   []interface{}
	mutex  *gmutex.Mutex
	buffer chan *Event
	err    error
}

func newEventSource(client *Client, method string, url string, data ...interface{}) *internalEventSource {
	s := &internalEventSource{
		client: client,
		method: method,
		url:    url,
		data:   data,
		mutex:  &gmutex.Mutex{},
		buffer: make(chan *Event, 1024),
	}
	g.Go(context.Background(), func(ctx context.Context) {
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
		defer s.close(scanner.Err())
		for s.processNextEvent(scanner) {
		}
	}, s.client.deferLogError)
	return s
}

func (s *internalEventSource) Event() <-chan *Event {
	return s.buffer
}

func (s *internalEventSource) Err() (err error) {
	s.mutex.LockFunc(func() {
		err = s.err
	})
	return
}

func (s *internalEventSource) Close() {
	for range s.buffer {
		// Drain the buffer
	}
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
			event.Id = strings.TrimLeft(strings.TrimPrefix(line, "id:"), " ")
		case strings.HasPrefix(line, "event:"):
			event.Event = strings.TrimLeft(strings.TrimPrefix(line, "event:"), " ")
		case strings.HasPrefix(line, "data:"):
			if event.Data != "" {
				event.Data += "\n"
			}
			event.Data += strings.TrimLeft(strings.TrimPrefix(line, "data:"), " ")
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
