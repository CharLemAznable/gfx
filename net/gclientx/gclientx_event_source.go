package gclientx

type Event struct {
	Id    string
	Event string
	Data  string
}

type EventSource interface {
	Event() <-chan Event
	Err() error
	Close()
}

func (c *Client) EventSource(method string, url string, data ...interface{}) EventSource {
	return newEventSource(c, method, url, data...)
}
