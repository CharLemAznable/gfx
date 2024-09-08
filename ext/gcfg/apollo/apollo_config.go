package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcmd"
)

type Config struct {
	Key   string `json:"key"`
	Watch bool   `json:"watch"`
}

const (
	defaultKey = "config"

	apolloKeyPattern   = "gf.gcfg.apollo.key"
	apolloWatchPattern = "gf.gcfg.apollo.watch"
)

func LoadConfig(ctx context.Context) (*Config, error) {
	fileConfigMap, err := agollox.LoadFileConfigMap(ctx)
	if err != nil {
		return nil, err
	}
	configMap := agollox.MapOmitNil(map[string]interface{}{
		"key":   gcmd.GetOptWithEnv(apolloKeyPattern).Val(),
		"watch": gcmd.GetOptWithEnv(apolloWatchPattern).Val(),
	})
	agollox.MapFill(configMap, agollox.MapOmitNil(map[string]interface{}{
		"key":   fileConfigMap["key"],
		"watch": fileConfigMap["watch"],
	}))
	config := &Config{Key: defaultKey, Watch: true}
	if err = gvar.New(configMap).Struct(config); err != nil {
		return nil, err
	}
	if config.Key == "" {
		config.Key = defaultKey
	}
	return config, nil
}
