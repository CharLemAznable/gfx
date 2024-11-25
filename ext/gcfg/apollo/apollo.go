package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

func LoadAdapter(ctx context.Context) (gcfg.Adapter, error) {
	return NewAdapterApollo(ctx)
}

func NewAdapterApollo(ctx context.Context) (adapter *AdapterApollo, err error) {
	config, err := LoadConfig(ctx)
	if err != nil {
		return
	}
	client, err := agollox.NewClient(ctx)
	if err != nil {
		return
	}
	adapter = &AdapterApollo{
		client: client,
		config: config,
		value:  gvar.New(nil, true),
		mutex:  &gmutex.Mutex{},
	}
	if config.Watch {
		client.SetChangeListener(adapter)
	}
	return
}

type AdapterApollo struct {
	client *agollox.Client
	config *Config
	value  *gvar.Var
	mutex  *gmutex.Mutex
}

func (a *AdapterApollo) Available(_ context.Context, _ ...string) bool {
	return a.client.Contains(a.config.Key)
}

func (a *AdapterApollo) Get(_ context.Context, pattern string) (value interface{}, err error) {
	if err = a.updateLocalValue(true); err != nil {
		return nil, err
	}
	return a.value.Val().(*gjson.Json).Get(pattern).Val(), nil
}

func (a *AdapterApollo) Data(_ context.Context) (data map[string]interface{}, err error) {
	if err = a.updateLocalValue(true); err != nil {
		return nil, err
	}
	return a.value.Val().(*gjson.Json).Map(), nil
}

func (a *AdapterApollo) OnChange(event *agollox.ChangeEvent) {
	if _, ok := event.Changes[a.config.Key]; ok {
		_ = a.updateLocalValue(false)
	}
}

func (a *AdapterApollo) updateLocalValue(onlyIfValueIsNil bool) (err error) {
	if !a.value.IsNil() && onlyIfValueIsNil {
		return
	}
	a.mutex.LockFunc(func() {
		if !a.value.IsNil() && onlyIfValueIsNil {
			return
		}
		var (
			value = gconv.Bytes(a.client.Get(a.config.Key))
			json  = gjson.New(nil)
		)
		if len(value) > 0 {
			json, err = gjson.LoadContent(value, true)
		}
		a.value.Set(json)
	})
	return
}
