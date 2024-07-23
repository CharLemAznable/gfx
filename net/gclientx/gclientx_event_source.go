package gclientx

type Event struct {
	Id    string
	Event string
	Data  string
}

type EventListener interface {
	OnEvent(event *Event, err error)
}

type EventListenerFunc func(event *Event, err error)

func (f EventListenerFunc) OnEvent(event *Event, err error) {
	f(event, err)
}

type EventSource interface {
	Execute(listener ...EventListener) (EventSource, error)
	Event() <-chan *Event
	Err() error
}

func (c *Client) EventSource(method string, url string, data ...interface{}) EventSource {
	return newEventSource(c, method, url, data...)
}
