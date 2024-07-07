package apollo

import (
	"context"
	"github.com/apolloconfig/agollo/v4"
	apolloConfig "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

type Config struct {
	AppID             string `v:"required"` // See apolloConfig.Config.
	IP                string `v:"required"` // See apolloConfig.Config.
	Cluster           string // See apolloConfig.Config. default: "default".
	NamespaceName     string // See apolloConfig.Config. default: "application".
	Key               string `v:"required"`
	IsBackupConfig    bool   // See apolloConfig.Config.
	BackupConfigPath  string // See apolloConfig.Config.
	Secret            string // See apolloConfig.Config.
	SyncServerTimeout int    // See apolloConfig.Config.
	MustStart         bool   // See apolloConfig.Config.
	Watch             bool   // Watch watches remote configuration updates, which updates local configuration in memory immediately when remote configuration changes.
}

type AdapterApollo struct {
	config Config        // Config object when created.
	client agollo.Client // Apollo client.
	value  *g.Var        // Configmap content cached. It is `*gjson.Json` value internally.
}

const (
	defaultCluster   = "default"
	defaultNamespace = "application"
)

func NewAdapterApollo(ctx context.Context, config Config) (adapter gcfg.Adapter, err error) {
	// Data validation.
	err = g.Validator().Data(config).Run(ctx)
	if err != nil {
		return nil, err
	}
	if config.Cluster == "" {
		config.Cluster = defaultCluster
	}
	if config.NamespaceName == "" {
		config.NamespaceName = defaultNamespace
	}
	client := &AdapterApollo{
		config: config,
		value:  g.NewVar(nil, true),
	}
	// Apollo client.
	client.client, err = agollo.StartWithConfig(func() (*apolloConfig.AppConfig, error) {
		return &apolloConfig.AppConfig{
			AppID:             config.AppID,
			Cluster:           config.Cluster,
			NamespaceName:     config.NamespaceName,
			IP:                config.IP,
			IsBackupConfig:    config.IsBackupConfig,
			BackupConfigPath:  config.BackupConfigPath,
			Secret:            config.Secret,
			SyncServerTimeout: config.SyncServerTimeout,
			MustStart:         config.MustStart,
		}, nil
	})
	if err != nil {
		return nil, gerror.Wrapf(err, `create apollo client failed with config: %+v`, config)
	}
	if config.Watch {
		client.client.AddChangeListener(client)
	}
	return client, nil
}

func (c *AdapterApollo) Available(_ context.Context, resource ...string) (ok bool) {
	if len(resource) == 0 && !c.value.IsNil() {
		return true
	}
	var namespace = c.config.NamespaceName
	if len(resource) > 0 && resource[0] != "" {
		namespace = resource[0]
	}
	return c.client.GetConfig(namespace) != nil
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

func (c *AdapterApollo) OnChange(event *storage.ChangeEvent) {
	if _, ok := event.Changes[c.config.Key]; ok {
		_ = c.updateLocalValue()
	}
}

func (c *AdapterApollo) OnNewestChange(_ *storage.FullChangeEvent) {
	// Nothing to do.
}

func (c *AdapterApollo) updateLocalValue() (err error) {
	cache := c.client.GetConfigCache(c.config.NamespaceName)
	defer cache.Clear()

	json := gjson.New(nil)
	content, err := cache.Get(c.config.Key)
	if err == nil {
		json, err = gjson.LoadContent(content, true)
		if err == nil {
			c.value.Set(json)
			return
		}
	}
	c.value.Set(json)
	return
}
