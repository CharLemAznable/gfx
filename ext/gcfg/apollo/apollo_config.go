package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/os/gcmd"
)

type Config struct {
	agollox.Config
	Key string `json:"key"`
}

const (
	defaultConfigFileName = "apollo"

	apolloAppIdPattern     = "gf.gcfg.apollo.appid"
	apolloClusterPattern   = "gf.gcfg.apollo.cluster"
	apolloNamespacePattern = "gf.gcfg.apollo.namespace"
	apolloIPPattern        = "gf.gcfg.apollo.ip"
	apolloKeyPattern       = "gf.gcfg.apollo.key"
)

func LoadConfig(ctx context.Context, fileName ...string) (*Config, error) {
	configFileName := defaultConfigFileName
	if len(fileName) > 0 && fileName[0] != "" {
		configFileName = fileName[0]
	}
	apolloConfig := &Config{Config: *agollox.DefaultConfig()}
	err := agollox.LoadConfig(ctx, apolloConfig, configFileName, map[string]interface{}{
		agollox.ConfigAppIdKey:     gcmd.GetOptWithEnv(apolloAppIdPattern).Val(),
		agollox.ConfigClusterKey:   gcmd.GetOptWithEnv(apolloClusterPattern).Val(),
		agollox.ConfigNamespaceKey: gcmd.GetOptWithEnv(apolloNamespacePattern).Val(),
		agollox.ConfigIPKey:        gcmd.GetOptWithEnv(apolloIPPattern).Val(),
		"key":                      gcmd.GetOptWithEnv(apolloKeyPattern).Val(),
	})
	if err != nil {
		return nil, err
	}
	return apolloConfig, nil
}
