package agollox

import (
	"context"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
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
)

func LoadConfig(ctx context.Context, pointer interface{}, fileName string, def ...map[string]interface{}) error {
	configFileMap, err := loadConfigFileMap(ctx, fileName)
	if err != nil {
		return err
	}
	configMap := loadDefMap(def...)
	for key, value := range configFileMap {
		if !gutil.MapContainsPossibleKey(configMap, key) {
			configMap[key] = value // fill config map with config file map
		}
	}
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
	return gutil.MapCopy(configMap), nil
}

func loadDefMap(def ...map[string]interface{}) map[string]interface{} {
	if len(def) > 0 && len(def[0]) > 0 {
		m := gutil.MapCopy(def[0])
		for k, v := range m {
			if v == nil {
				delete(m, k)
			}
		}
		return m
	}
	return map[string]interface{}{}
}
