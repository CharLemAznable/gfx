package gcfgx_test

import (
	"context"
	"github.com/CharLemAznable/gfx/errors/gerrorx"
	"github.com/CharLemAznable/gfx/os/gcfgx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

var (
	ctx = context.TODO()

	normalCfgx = func() *gcfgx.Config {
		c := gcfgx.Instance("normal")
		c.SetAdapter(&normalAdapter{})
		return c
	}()

	errorCfgx = func() *gcfgx.Config {
		c := gcfgx.Instance("error")
		c.SetAdapter(&errorAdapter{})
		return c
	}()

	testErr = gerrorx.ErrorString("error")
)

func Test_Default(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		normalCfgx.SetErrorFn(nil)
		t.Assert(normalCfgx.MustGet(ctx, "key1", "value1"), "[key1]")
		t.Assert(normalCfgx.MustGetWithEnv(ctx, "key2", "value2"), "[key2]")
		t.Assert(normalCfgx.MustGetWithCmd(ctx, "key3", "value3"), "[key3]")
		t.Assert(normalCfgx.MustData(ctx)["key"], "value")
		t.Assert(normalCfgx.MustGetWithCmdAndEnv(ctx, "key3", "value3"), "[key3]")

		errorCfgx.SetErrorFn(nil)
		t.Assert(errorCfgx.MustGet(ctx, "key1", "value1"), "value1")
		t.Assert(errorCfgx.MustGetWithEnv(ctx, "key2", "value2"), "value2")
		t.Assert(errorCfgx.MustGetWithCmd(ctx, "key3", "value3"), "value3")
		t.AssertNil(errorCfgx.MustData(ctx))
		t.Assert(errorCfgx.MustGetWithCmdAndEnv(ctx, "key3", "value3"), "value3")
	})
}

func Test_SetErrorFn(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ffn := func(ctx context.Context, format string, v ...interface{}) {
			t.Assert(v[0], testErr)
		}

		normalCfgx.SetErrorFn(ffn)
		t.Assert(normalCfgx.MustGet(ctx, "key1", "value1"), "[key1]")
		t.Assert(normalCfgx.MustGetWithEnv(ctx, "key2", "value2"), "[key2]")
		t.Assert(normalCfgx.MustGetWithCmd(ctx, "key3", "value3"), "[key3]")
		t.Assert(normalCfgx.MustData(ctx)["key"], "value")
		t.Assert(normalCfgx.MustGetWithCmdAndEnv(ctx, "key3", "value3"), "[key3]")

		errorCfgx.SetErrorFn(ffn)
		t.AssertNil(errorCfgx.MustGet(ctx, "key1", "value1"))
		t.AssertNil(errorCfgx.MustGetWithEnv(ctx, "key2", "value2"))
		t.AssertNil(errorCfgx.MustGetWithCmd(ctx, "key3", "value3"))
		t.AssertNil(errorCfgx.MustData(ctx))
		t.AssertNil(errorCfgx.MustGetWithCmdAndEnv(ctx, "key3", "value3"))
	})
}

func Test_SetErrorLogger(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gcfgx.Instance().SetErrorLogger(g.Log()), nil)
		t.AssertNE(gcfgx.Instance().SetErrorLogger(nil), nil)
	})
}

func Test_GetWithCmdAndEnv(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(normalCfgx.MustGetWithCmdAndEnv(ctx, "not.found", "def"), "def")

		_ = genv.Set("NOT_FOUND", "DEFAULT")
		defer func() { _ = genv.Remove("NOT_FOUND") }()
		t.Assert(normalCfgx.MustGetWithCmdAndEnv(ctx, "not.found", "def"), "DEFAULT")
	})
}

type normalAdapter struct{}

func (a *normalAdapter) Available(_ context.Context, _ ...string) (ok bool) {
	return true
}

func (a *normalAdapter) Get(_ context.Context, pattern string) (value interface{}, err error) {
	if pattern == "not.found" {
		return nil, nil
	}
	return "[" + pattern + "]", nil
}

func (a *normalAdapter) Data(_ context.Context) (data map[string]interface{}, err error) {
	return g.Map{"key": "value"}, nil
}

type errorAdapter struct{}

func (a *errorAdapter) Available(_ context.Context, _ ...string) (ok bool) {
	return false
}

func (a *errorAdapter) Get(_ context.Context, _ string) (value interface{}, err error) {
	return nil, testErr
}

func (a *errorAdapter) Data(_ context.Context) (data map[string]interface{}, err error) {
	return nil, testErr
}
