package agollox

import (
	"context"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gutil"
)

type Config = config.AppConfig

const (
	defaultCluster   = "default"
	defaultNamespace = "application"
)

func DefaultConfig() *Config {
	return &config.AppConfig{
		Cluster:        defaultCluster,
		NamespaceName:  defaultNamespace,
		IsBackupConfig: true,
	}
}

//goland:noinspection GoUnusedConst
const (
	ConfigAppIdKey             = "appId"
	ConfigClusterKey           = "cluster"
	ConfigNamespaceKey         = "namespaceName"
	ConfigIPKey                = "ip"
	ConfigIsBackupConfigKey    = "isBackupConfig"
	ConfigBackupConfigPathKey  = "backupConfigPath"
	ConfigSecretKey            = "secret"
	ConfigLabelKey             = "label"
	ConfigSyncServerTimeoutKey = "syncServerTimeout"

	apolloAppIdPattern     = "gf.apollo.appid"
	apolloClusterPattern   = "gf.apollo.cluster"
	apolloNamespacePattern = "gf.apollo.namespace"
	apolloIPPattern        = "gf.apollo.ip"
)

func LoadConfig(ctx context.Context, pointer interface{}, fileName string, def ...map[string]interface{}) error {
	configFileMap, err := loadConfigFileMap(ctx, fileName)
	if err != nil {
		return err
	}
	configMap := loadDefMap(def...)
	mapFill(configMap, configFileMap)
	return gvar.New(configMap).Struct(pointer)
}

func loadConfigFileMap(ctx context.Context, fileName string) (map[string]interface{}, error) {
	configFile, err := gcfg.NewAdapterFile(fileName)
	if err != nil {
		return nil, err
	}
	if !configFile.Available(ctx) {
		return map[string]interface{}{}, nil
	}
	configMap, err := configFile.Data(ctx)
	if err != nil {
		return nil, err
	}
	return mapOmitNil(gutil.MapCopy(configMap)), nil
}

func loadDefMap(def ...map[string]interface{}) map[string]interface{} {
	cmdMap := mapOmitNil(map[string]interface{}{
		ConfigAppIdKey:     gcmd.GetOptWithEnv(apolloAppIdPattern).Val(),
		ConfigClusterKey:   gcmd.GetOptWithEnv(apolloClusterPattern).Val(),
		ConfigNamespaceKey: gcmd.GetOptWithEnv(apolloNamespacePattern).Val(),
		ConfigIPKey:        gcmd.GetOptWithEnv(apolloIPPattern).Val(),
	})
	if len(def) > 0 && len(def[0]) > 0 {
		defMap := mapOmitNil(gutil.MapCopy(def[0]))
		mapFill(defMap, cmdMap)
		return defMap
	}
	return cmdMap
}

func mapOmitNil(data map[string]interface{}) map[string]interface{} {
	if len(data) == 0 {
		return data
	}
	for k, v := range data {
		if v == nil {
			delete(data, k)
		}
	}
	return data
}

func mapFill(dst map[string]interface{}, src map[string]interface{}) {
	for key, value := range src {
		if !gutil.MapContainsPossibleKey(dst, key) {
			dst[key] = value // fill dst map with src map
		}
	}
}
