package gcmdx

import (
	"github.com/CharLemAznable/gfx/internal/utils"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcmd"
	"os"
	"strings"
)

func GetOpt(name string, def ...string) *gvar.Var {
	if v := gcmd.GetOpt(name, def...); !v.IsEmpty() {
		return v
	}
	return utils.DefaultOrNil[string](def...)
}

func GetOptWithEnv(key string, def ...interface{}) *gvar.Var {
	cmdKey := formatCmdKey(key)
	if v := gcmd.GetOpt(cmdKey); !v.IsEmpty() {
		return v
	}
	envKey := formatEnvKey(key)
	if r, ok := os.LookupEnv(envKey); ok && r != "" {
		return gvar.New(r)
	}
	return utils.DefaultOrNil[any](def...)
}

func formatCmdKey(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", "."))
}

func formatEnvKey(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, ".", "_"))
}
