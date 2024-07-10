package apollo

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
)

const (
	defaultApolloConfigFileName = "apollo"
	apolloConfigKey             = "apollo"

	apolloAppIdPattern     = "gf.gcfg.apollo.appid"
	apolloIPPattern        = "gf.gcfg.apollo.ip"
	apolloClusterPattern   = "gf.gcfg.apollo.cluster"
	apolloNamespacePattern = "gf.gcfg.apollo.namespace"
	apolloKeyPattern       = "gf.gcfg.apollo.key"
)

func LoadAdapter(ctx context.Context, fileName ...string) (gcfg.Adapter, error) {
	apolloConfigFileName := defaultApolloConfigFileName
	if len(fileName) > 0 && fileName[0] != "" {
		apolloConfigFileName = fileName[0]
	}
	apolloCfg, _ := g.Cfg(apolloConfigFileName).Get(ctx, apolloConfigKey)
	apolloConfig := &Config{}
	_ = apolloCfg.Struct(apolloConfig)

	if apolloConfig.AppID == "" {
		apolloConfig.AppID = gcmd.GetOptWithEnv(apolloAppIdPattern).String()
	}
	if apolloConfig.IP == "" {
		apolloConfig.IP = gcmd.GetOptWithEnv(apolloIPPattern).String()
	}
	if apolloConfig.Cluster == "" {
		apolloConfig.Cluster = gcmd.GetOptWithEnv(apolloClusterPattern).String()
	}
	if apolloConfig.NamespaceName == "" {
		apolloConfig.NamespaceName = gcmd.GetOptWithEnv(apolloNamespacePattern).String()
	}
	if apolloConfig.Key == "" {
		apolloConfig.Key = gcmd.GetOptWithEnv(apolloKeyPattern).String()
	}
	return NewAdapterApollo(ctx, apolloConfig)
}
