package apollo

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"os"
	"strings"
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
	apolloConfig := Config{}
	_ = apolloCfg.Struct(&apolloConfig)

	if apolloConfig.AppID == "" {
		apolloConfig.AppID = getOptWithEnv(apolloAppIdPattern)
	}
	if apolloConfig.IP == "" {
		apolloConfig.IP = getOptWithEnv(apolloIPPattern)
	}
	if apolloConfig.Cluster == "" {
		apolloConfig.Cluster = getOptWithEnv(apolloClusterPattern)
	}
	if apolloConfig.NamespaceName == "" {
		apolloConfig.NamespaceName = getOptWithEnv(apolloNamespacePattern)
	}
	if apolloConfig.Key == "" {
		apolloConfig.Key = getOptWithEnv(apolloKeyPattern)
	}
	return NewAdapterApollo(ctx, apolloConfig)
}

func getOptWithEnv(key string) string {
	if v := gcmd.GetOpt(key); !v.IsEmpty() {
		return v.String()
	}
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if r, ok := os.LookupEnv(envKey); ok && r != "" {
		return r
	}
	return ""
}
