package gclientx

type Event struct {
	Id    string
	Event string
	Data  string
}

type EventListener interface {
	OnEvent(event *Event)
	OnClose(err error)
}

type EventListenerFunc func(event *Event, err error)

func (f EventListenerFunc) OnEvent(event *Event) {
	f(event, nil)
}

func (f EventListenerFunc) OnClose(err error) {
	f(nil, err)
}

type EventListenerChan chan<- *Event

func (c EventListenerChan) OnEvent(event *Event) {
	c <- event
}

func (c EventListenerChan) OnClose(_ error) {
	close(c)
}

type EventSource interface {
	Execute(listener ...EventListener) EventSource
	Done() <-chan struct{}
	Err() error
}

func (c *Client) EventSource(method string, url string, data ...interface{}) EventSource {
	return newEventSource(c, method, url, data...)
}
