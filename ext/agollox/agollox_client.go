package agollox

import (
	"context"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

type Client struct {
	client   *configClient
	mapValue *gvar.Var
	mapMutex *gmutex.Mutex
	listener *gvar.Var
}

func NewClient(ctx context.Context) (*Client, error) {
	intClient, err := agolloInstance(ctx) // agollo client 为单例
	if err != nil {
		return nil, gerror.Wrapf(err, `create agollo client failed`)
	}
	client := &Client{
		client:   intClient,
		mapValue: gvar.New(nil, true),
		mapMutex: &gmutex.Mutex{},
		listener: gvar.New(nil, true),
	}
	intClient.client.AddChangeListener(client)
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
		g.Go(context.Background(), func(_ context.Context) {
			listener.OnChange(event)
		}, nil)
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
		intClient := c.client
		config.SplitNamespaces(intClient.config.NamespaceName, func(namespace string) {
			// 如果配置了多namespace, 会合并所有的键值对
			// 且后置的namespace会覆盖前置的namespace的同名键
			cache := intClient.client.GetConfigCache(namespace)
			if cache != nil {
				cache.Range(func(key, value interface{}) bool {
					m.Set(gconv.String(key), value)
					return true
				})
			}
		})
		c.mapValue.Set(m)
	})
}

type configClient struct {
	config *Config
	client agollo.Client
}

var (
	agolloVar   = gvar.New(nil, true)
	agolloMutex = &gmutex.Mutex{}
)

func agolloInstance(ctx context.Context) (client *configClient, err error) {
	if !agolloVar.IsNil() {
		client = agolloVar.Val().(*configClient)
		return
	}
	agolloMutex.LockFunc(func() {
		if !agolloVar.IsNil() {
			client = agolloVar.Val().(*configClient)
			return
		}
		var (
			cfg *Config
			cli agollo.Client
		)
		cfg, err = LoadConfig(ctx)
		if err != nil {
			return
		}
		cli, err = agollo.StartWithConfig(func() (*Config, error) {
			return cfg, nil
		})
		if err != nil {
			return
		}
		client = &configClient{
			config: cfg,
			client: cli,
		}
		agolloVar.Set(client)
	})
	return
}
