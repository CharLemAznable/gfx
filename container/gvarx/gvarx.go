package gvarx

import (
	"github.com/CharLemAznable/gfx/internal/utils"
	"github.com/gogf/gf/v2/container/gvar"
)

func DefaultOrNil(def ...interface{}) *gvar.Var {
	return utils.DefaultOrNil[any](def...)
}
