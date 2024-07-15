package agollox

import "github.com/apolloconfig/agollo/v4/storage"

type ChangeEvent = storage.ChangeEvent

type ChangeListener interface {
	OnChange(event *ChangeEvent)
}

type ChangeListenerFunc func(event *ChangeEvent)

func (f ChangeListenerFunc) OnChange(event *ChangeEvent) {
	f(event)
}
