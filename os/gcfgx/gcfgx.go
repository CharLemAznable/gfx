package gcfgx

import (
	"context"
	"github.com/CharLemAznable/gfx/container/gvarx"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
)

type Config struct {
	config  *gcfg.Config
	errorFn func(ctx context.Context, format string, v ...interface{})
}

func New(config *gcfg.Config) *Config {
	return &Config{config: config}
}

func (c *Config) SetErrorFn(errorFn func(ctx context.Context, format string, v ...interface{})) *Config {
	c.errorFn = errorFn
	return c
}

func (c *Config) SetErrorLogger(logger *glog.Logger) *Config {
	if logger != nil {
		return c.SetErrorFn(logger.Fatalf)
	} else {
		return c.SetErrorFn(nil)
	}
}

func (c *Config) MustGet(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {
	v, err := c.config.Get(ctx, pattern, def...)
	if err != nil {
		if c.errorFn != nil {
			c.errorFn(ctx, `%+v`, err)
		} else {
			return gvarx.DefaultOrNil(def...)
		}
	}
	return v
}

func (c *Config) MustGetWithEnv(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {
	v, err := c.config.GetWithEnv(ctx, pattern, def...)
	if err != nil {
		if c.errorFn != nil {
			c.errorFn(ctx, `%+v`, err)
		} else {
			return gvarx.DefaultOrNil(def...)
		}
	}
	return v
}

func (c *Config) MustGetWithCmd(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {
	v, err := c.config.GetWithCmd(ctx, pattern, def...)
	if err != nil {
		if c.errorFn != nil {
			c.errorFn(ctx, `%+v`, err)
		} else {
			return gvarx.DefaultOrNil(def...)
		}
	}
	return v
}

func (c *Config) MustData(ctx context.Context) map[string]interface{} {
	v, err := c.config.Data(ctx)
	if err != nil && c.errorFn != nil {
		c.errorFn(ctx, `%+v`, err)
	}
	return v
}
