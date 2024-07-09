package gcmdx

import (
	"github.com/CharLemAznable/gfx/internal/utils"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcmd"
	"os"
)

func GetOpt(name string, def ...string) *gvar.Var {
	cmdKey := utils.FormatCmdKey(name)
	if v := gcmd.GetOpt(cmdKey, def...); !v.IsEmpty() {
		return v
	}
	return utils.DefaultOrNil[string](def...)
}

func GetOptWithEnv(key string, def ...interface{}) *gvar.Var {
	cmdKey := utils.FormatCmdKey(key)
	if v := gcmd.GetOpt(cmdKey); !v.IsEmpty() {
		return v
	}
	envKey := utils.FormatEnvKey(key)
	if r, ok := os.LookupEnv(envKey); ok && r != "" {
		return gvar.New(r)
	}
	return utils.DefaultOrNil[any](def...)
}
