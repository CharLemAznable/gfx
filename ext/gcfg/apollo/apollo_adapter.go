package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

type Config struct {
	agollox.Config
	Key string
}

type AdapterApollo struct {
	client *agollox.Client
	config *Config
	value  *gvar.Var
}

var (
	configRules = g.MapStrStr{
		"AppID": "required",
		"IP":    "required",
		"Key":   "required",
	}
)

func NewAdapterApollo(ctx context.Context, config *Config) (adapter gcfg.Adapter, err error) {
	// Data validation.
	err = g.Validator().Rules(configRules).Data(config).Run(ctx)
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
	}
	agolloClient.SetOnChangeFn(client.onChange)
	return client, nil
}

func (c *AdapterApollo) Available(_ context.Context, _ ...string) bool {
	return c.client.Available()
}

func (c *AdapterApollo) Get(_ context.Context, pattern string) (value interface{}, err error) {
	if c.value.IsNil() {
		if err = c.updateLocalValue(); err != nil {
			return nil, err
		}
	}
	return c.value.Val().(*gjson.Json).Get(pattern).Val(), nil
}

func (c *AdapterApollo) Data(_ context.Context) (data map[string]interface{}, err error) {
	if c.value.IsNil() {
		if err = c.updateLocalValue(); err != nil {
			return nil, err
		}
	}
	return c.value.Val().(*gjson.Json).Map(), nil
}

func (c *AdapterApollo) onChange(event *storage.ChangeEvent) {
	if _, ok := event.Changes[c.config.Key]; !ok {
		return
	}
	_ = c.updateLocalValue()
}

func (c *AdapterApollo) updateLocalValue() error {
	value := c.client.Get(c.config.Key)
	json, err := gjson.LoadContent(value, true)
	if err != nil {
		c.value.Set(gjson.New(nil))
		return err
	}
	c.value.Set(json)
	return nil
}
