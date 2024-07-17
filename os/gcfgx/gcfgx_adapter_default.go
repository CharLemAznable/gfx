package gcfgx

import (
	"context"
	"github.com/gogf/gf/v2/os/gcfg"
)

type AdapterDefault struct {
	adapter   gcfg.Adapter
	fallbacks []gcfg.Adapter
}

func NewAdapterDefault(adapter gcfg.Adapter, fallbacks ...gcfg.Adapter) *AdapterDefault {
	return &AdapterDefault{
		adapter:   adapter,
		fallbacks: fallbacks,
	}
}

func (a *AdapterDefault) Available(ctx context.Context, resource ...string) (ok bool) {
	if ok = a.adapter.Available(ctx, resource...); ok {
		return
	}
	for _, fallback := range a.fallbacks {
		if ok = fallback.Available(ctx, resource...); ok {
			return
		}
	}
	return
}

func (a *AdapterDefault) Get(ctx context.Context, pattern string) (value interface{}, err error) {
	if value, err = a.adapter.Get(ctx, pattern); err == nil && value != nil {
		return
	}
	for _, fallback := range a.fallbacks {
		if value, err = fallback.Get(ctx, pattern); err == nil && value != nil {
			return
		}
	}
	return nil, nil
}

func (a *AdapterDefault) Data(ctx context.Context) (data map[string]interface{}, err error) {
	if data, err = a.adapter.Data(ctx); err == nil && data != nil {
		return
	}
	for _, fallback := range a.fallbacks {
		if data, err = fallback.Data(ctx); err == nil && data != nil {
			return
		}
	}
	return nil, nil
}
