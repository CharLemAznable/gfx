package utils

import "github.com/gogf/gf/v2/container/gvar"

func DefaultOrNil[T any](def ...T) *gvar.Var {
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}
