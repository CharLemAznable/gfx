package gviewx

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gview"
)

type View struct {
	adapter Adapter
	view    *gview.View
}

type Adapter interface {
	GetContent(key string) (content string, err error)
}

func New() *View {
	return NewWithAdapter(NewAdapterFile())
}

func NewWithAdapter(adapter Adapter) *View {
	return &View{
		adapter: adapter,
		view:    gview.New(),
	}
}

const DefaultInstanceName = "default"

var localInstances = gmap.NewStrAnyMap(true)

func Instance(name ...string) *View {
	var instanceName = DefaultInstanceName
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	return localInstances.GetOrSetFuncLock(instanceName, func() interface{} {
		return New()
	}).(*View)
}

func (view *View) SetAdapter(adapter Adapter) *View {
	view.adapter = adapter
	return view
}

func (view *View) GetAdapter() Adapter {
	return view.adapter
}

func (view *View) Parse(ctx context.Context, key string, params ...gview.Params) (string, error) {
	content, err := view.adapter.GetContent(key)
	if err != nil {
		return "", err
	}
	return view.view.ParseContent(ctx, content, params...)
}
