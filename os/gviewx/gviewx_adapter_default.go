package gviewx

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type AdapterDefault struct {
	adapter   Adapter
	fallbacks []Adapter
}

func NewAdapterDefault(adapter Adapter, fallbacks ...Adapter) *AdapterDefault {
	return &AdapterDefault{
		adapter:   adapter,
		fallbacks: fallbacks,
	}
}

func (a *AdapterDefault) GetContent(key string) (content string, err error) {
	if content, err = a.adapter.GetContent(key); err == nil {
		return
	}
	for _, fallback := range a.fallbacks {
		if content, err = fallback.GetContent(key); err == nil {
			return
		}
	}
	return "", gerror.NewCodef(gcode.CodeInvalidParameter, `template "%s" not found`, key)
}
