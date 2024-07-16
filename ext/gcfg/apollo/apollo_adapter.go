package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcfg"
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
		"appId": "required",
		"ip":    "required",
		"key":   "required",
	}
	configMessage = map[string]interface{}{
		"appId": "Apollo AppId field is required",
		"ip":    "Apollo IP field is required",
		"key":   "Apollo Key field is required",
	}
)

func NewAdapterApollo(ctx context.Context, config *Config) (adapter gcfg.Adapter, err error) {
	// Data validation.
	err = gvalid.New().Rules(configRules).Messages(configMessage).Data(config).Run(ctx)
	if err != nil {
		return nil, err
	}
	agolloClient, err := agollox.NewClient(&config.Config)
	if err != nil {
		return nil, err
	}
	client := &AdapterApollo{
		client: agolloClient,
		config: config,
		value:  gvar.New(nil, true),
		mutex:  &gmutex.Mutex{},
	}
	if config.Watch {
		agolloClient.SetChangeListener(client)
	}
	return client, nil
}

func (c *AdapterApollo) Available(_ context.Context, _ ...string) bool {
	return c.client.Contains(c.config.Key)
}

func (c *AdapterApollo) Get(_ context.Context, pattern string) (value interface{}, err error) {
	if err = c.updateLocalValue(false); err != nil {
		return nil, err
	}
	return c.value.Val().(*gjson.Json).Get(pattern).Val(), nil
}

func (c *AdapterApollo) Data(_ context.Context) (data map[string]interface{}, err error) {
	if err = c.updateLocalValue(false); err != nil {
		return nil, err
	}
	return c.value.Val().(*gjson.Json).Map(), nil
}

func (c *AdapterApollo) OnChange(event *agollox.ChangeEvent) {
	if _, ok := event.Changes[c.config.Key]; ok {
		_ = c.updateLocalValue(true)
	}
}

func (c *AdapterApollo) updateLocalValue(anyway bool) (err error) {
	if !(c.value.IsNil() || anyway) {
		return
	}
	c.mutex.LockFunc(func() {
		if !(c.value.IsNil() || anyway) {
			return
		}
		var (
			value = c.client.Get(c.config.Key)
			json  = gjson.New(nil)
		)
		json, err = gjson.LoadContent(value, true)
		c.value.Set(json)
	})
	return
}
