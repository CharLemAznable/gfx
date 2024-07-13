package apollo

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/agollox"
	"github.com/gogf/gf/v2/os/gcmd"
)

type Config = agollox.Config

const (
	defaultConfigFileName = "apollo"

	apolloAppIdPattern     = "gf.gviewx.apollo.appid"
	apolloClusterPattern   = "gf.gviewx.apollo.cluster"
	apolloNamespacePattern = "gf.gviewx.apollo.namespace"
	apolloIPPattern        = "gf.gviewx.apollo.ip"
)

func LoadConfig(ctx context.Context, fileName ...string) (*Config, error) {
	configFileName := defaultConfigFileName
	if len(fileName) > 0 && fileName[0] != "" {
		configFileName = fileName[0]
	}
	apolloConfig := agollox.DefaultConfig()
	err := agollox.LoadConfig(ctx, apolloConfig, configFileName, map[string]interface{}{
		agollox.ConfigAppIdKey:     gcmd.GetOptWithEnv(apolloAppIdPattern).Val(),
		agollox.ConfigClusterKey:   gcmd.GetOptWithEnv(apolloClusterPattern).Val(),
		agollox.ConfigNamespaceKey: gcmd.GetOptWithEnv(apolloNamespacePattern).Val(),
		agollox.ConfigIPKey:        gcmd.GetOptWithEnv(apolloIPPattern).Val(),
	})
	if err != nil {
		return nil, err
	}
	return apolloConfig, nil
}
