package gcfgx

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gcfg"
)

var localInstances = gmap.NewStrAnyMap(true)

func Instance(name ...string) *Config {
	var instanceName = gcfg.DefaultInstanceName
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	return localInstances.GetOrSetFuncLock(instanceName, func() interface{} {
		return New(gcfg.Instance(instanceName))
	}).(*Config)
}
