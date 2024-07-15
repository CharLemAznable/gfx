package agollox

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type ChangeEvent = storage.ChangeEvent

type Client struct {
	client      agollo.Client
	config      *Config
	mapping     *gmap.StrAnyMap
	initialized bool
	onChangeFn  func(event *ChangeEvent)
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
		client:  agolloClient,
		config:  config,
		mapping: gmap.NewStrAnyMap(true),
	}
	client.client.AddChangeListener(client)
	return client, nil
}

func (c *Client) Contains(key string) bool {
	c.initialize()
	return c.mapping.Contains(key)
}

func (c *Client) Get(key string) interface{} {
	c.initialize()
	return c.mapping.Get(key)
}

func (c *Client) Map() map[string]interface{} {
	c.initialize()
	return c.mapping.Map()
}

func (c *Client) SetOnChangeFn(fn func(event *storage.ChangeEvent)) *Client {
	c.onChangeFn = fn
	return c
}

func (c *Client) OnChange(event *storage.ChangeEvent) {
	c.updateLocalMapping()
	if c.onChangeFn != nil {
		c.onChangeFn(event)
	}
}

func (c *Client) OnNewestChange(_ *storage.FullChangeEvent) {
	// Nothing to do.
}

func (c *Client) initialize() {
	if !c.initialized {
		c.updateLocalMapping()
	}
}

func (c *Client) updateLocalMapping() {
	cache := c.client.GetConfigCache(c.config.NamespaceName)
	defer cache.Clear()
	cache.Range(func(key, value interface{}) bool {
		c.mapping.Set(gconv.String(key), value)
		return true
	})
	c.initialized = true
}
