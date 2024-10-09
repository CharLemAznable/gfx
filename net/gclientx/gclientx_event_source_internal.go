package gclientx

import (
	"bufio"
	"github.com/gogf/gf/v2/os/gmutex"
	"strings"
)

type internalEventSource struct {
	mutex  *gmutex.Mutex
	buffer chan *Event
	err    error
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
