package agollox

import (
	"context"
	agolloConfig "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/gogf/gf/v2/util/gvalid"
)

type Config = agolloConfig.AppConfig

const (
	defaultCluster          = "default"
	defaultNamespace        = "application"
	defaultBackupConfigPath = ".apollo.bk"
)

func DefaultConfig() *Config {
	return &agolloConfig.AppConfig{
		Cluster:          defaultCluster,
		NamespaceName:    defaultNamespace,
		IsBackupConfig:   true,
		BackupConfigPath: defaultBackupConfigPath,
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
	ConfigMustStart            = "mustStart"

	apolloAppIdPattern     = "gf.apollo.appid"
	apolloClusterPattern   = "gf.apollo.cluster"
	apolloNamespacePattern = "gf.apollo.namespace"
	apolloIPPattern        = "gf.apollo.ip"
)

const (
	defaultConfigFileName = "apollo"

	apolloConfigFileNamePattern = "gf.apollo.config.file"
)

func LoadFileConfigMap(ctx context.Context) (map[string]interface{}, error) {
	fileName := defaultConfigFileName
	if optFileName := gcmd.GetOptWithEnv(apolloConfigFileNamePattern).String(); optFileName != "" {
		fileName = optFileName
	}
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
	return MapOmitNil(gutil.MapCopy(configMap)), nil
}

var (
	configRules = map[string]string{
		"appId": "required",
		"ip":    "required",
	}
	configMessage = map[string]interface{}{
		"appId": "Apollo AppId field is required",
		"ip":    "Apollo IP field is required",
	}
)

func LoadConfig(ctx context.Context) (*Config, error) {
	fileConfigMap, err := LoadFileConfigMap(ctx)
	if err != nil {
		return nil, err
	}
	configMap := MapOmitNil(map[string]interface{}{
		ConfigAppIdKey:     gcmd.GetOptWithEnv(apolloAppIdPattern).Val(),
		ConfigClusterKey:   gcmd.GetOptWithEnv(apolloClusterPattern).Val(),
		ConfigNamespaceKey: gcmd.GetOptWithEnv(apolloNamespacePattern).Val(),
		ConfigIPKey:        gcmd.GetOptWithEnv(apolloIPPattern).Val(),
	})
	MapFill(configMap, fileConfigMap)
	config := DefaultConfig()
	if err = gvar.New(configMap).Struct(config); err != nil {
		return nil, err
	}
	if config.Cluster == "" {
		config.Cluster = defaultCluster
	}
	if config.NamespaceName == "" {
		config.NamespaceName = defaultNamespace
	}
	if config.BackupConfigPath == "" {
		config.BackupConfigPath = defaultBackupConfigPath
	}
	if err = gvalid.New().
		Rules(configRules).Messages(configMessage).
		Data(config).Run(ctx); err != nil {
		return nil, err
	}
	return config, nil
}
