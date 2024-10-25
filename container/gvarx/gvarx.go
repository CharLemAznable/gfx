package gvarx

import (
	"github.com/CharLemAznable/gfx/internal/utils"
	"github.com/gogf/gf/v2/container/gvar"
)

func DefaultOrNil(def ...interface{}) *gvar.Var {
	return utils.DefaultOrNil(def...)
}

func DefaultIfNil(v *gvar.Var, def ...interface{}) *gvar.Var {
	if v == nil || v.IsNil() {
		return utils.DefaultOrNil(def...)
	}
	return v
}

func DefaultIfEmpty(v *gvar.Var, def ...interface{}) *gvar.Var {
	if v == nil || v.IsEmpty() {
		return utils.DefaultOrNil(def...)
	}
	return v
}
