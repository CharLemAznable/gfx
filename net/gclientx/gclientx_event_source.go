package gclientx

import "context"

type Event struct {
	Id    string
	Event string
	Data  string
}

type EventSource interface {
	Execute(ctx context.Context) error
	Event() <-chan *Event
	Err() error
}

func (c *Client) EventSource(method string, url string, data ...interface{}) EventSource {
	return newEventReader(c, method, url, data...)
}
