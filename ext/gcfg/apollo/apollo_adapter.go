package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/gogf/gf/v2/util/gvalid"
)

type AdapterApollo struct {
	client *agollox.Client
	config *Config
	value  *gvar.Var
	mutex  *gmutex.Mutex
}

var (
	configRules = map[string]string{
		"key": "required",
	}
	configMessage = map[string]interface{}{
		"key": "Apollo Key field is required",
	}
)

func NewAdapterApollo(ctx context.Context, config *Config) (adapter *AdapterApollo, err error) {
	// Data validation.
	err = gvalid.New().Rules(configRules).Messages(configMessage).Data(config).Run(ctx)
	if err != nil {
		return
	}
	agolloClient, err := agollox.NewClient(ctx, &config.Config)
	if err != nil {
		return
	}
	adapter = &AdapterApollo{
		client: agolloClient,
		config: config,
		value:  gvar.New(nil, true),
		mutex:  &gmutex.Mutex{},
	}
	if config.Watch {
		agolloClient.SetChangeListener(adapter)
	}
	return
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
			value = a.client.Get(a.config.Key)
			json  = gjson.New(nil)
		)
		json, err = gjson.LoadContent(value, true)
		a.value.Set(json)
	})
	return
}
