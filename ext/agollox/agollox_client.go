package agollox

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

type Client struct {
	client   agollo.Client
	config   *Config
	mapValue *gvar.Var
	mapMutex *gmutex.Mutex
	listener *gvar.Var
}

func NewClient(config *Config) (*Client, error) {
	if config.Cluster == "" {
		config.Cluster = defaultCluster
	}
	if config.NamespaceName == "" {
		config.NamespaceName = defaultNamespace
	}
	agolloClient, err := agollo.StartWithConfig(func() (*Config, error) {
		return config, nil
	})
	if err != nil {
		return nil, gerror.Wrapf(err, `create agollo client failed with config: %+v`, config)
	}
	client := &Client{
		client:   agolloClient,
		config:   config,
		mapValue: gvar.New(nil, true),
		mapMutex: &gmutex.Mutex{},
		listener: gvar.New(nil, true),
	}
	client.client.AddChangeListener(client)
	return client, nil
}

func (c *Client) Contains(key string) bool {
	c.updateLocalMapping(true)
	return c.mapValue.Val().(*gmap.StrAnyMap).Contains(key)
}

func (c *Client) Get(key string) interface{} {
	c.updateLocalMapping(true)
	return c.mapValue.Val().(*gmap.StrAnyMap).Get(key)
}

func (c *Client) Map() map[string]interface{} {
	c.updateLocalMapping(true)
	return c.mapValue.Val().(*gmap.StrAnyMap).Map()
}

func (c *Client) SetChangeListener(listener ChangeListener) *Client {
	c.listener.Set(listener)
	return c
}

func (c *Client) OnChange(event *storage.ChangeEvent) {
	c.updateLocalMapping(false)
	if listener, ok := c.listener.Val().(ChangeListener); ok && listener != nil {
		go listener.OnChange(event)
	}
}

func (c *Client) OnNewestChange(_ *storage.FullChangeEvent) {
	// Nothing to do.
}

func (c *Client) updateLocalMapping(onlyIfValueIsNil bool) {
	if !c.mapValue.IsNil() && onlyIfValueIsNil {
		return
	}
	c.mapMutex.LockFunc(func() {
		if !c.mapValue.IsNil() && onlyIfValueIsNil {
			return
		}
		m := gmap.NewStrAnyMap(true)
		cache := c.client.GetConfigCache(c.config.NamespaceName)
		cache.Range(func(key, value interface{}) bool {
			m.Set(gconv.String(key), value)
			return true
		})
		c.mapValue.Set(m)
	})
}
