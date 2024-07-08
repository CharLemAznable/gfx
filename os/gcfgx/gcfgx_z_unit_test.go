package gcfgx_test

import (
	"context"
	"github.com/CharLemAznable/gfx/errors/gerrorx"
	"github.com/CharLemAznable/gfx/os/gcfgx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

var (
	ctx = context.TODO()

	normalAdpt = &normalAdapter{}
	normalCfgx = func() *gcfgx.Config {
		c := gcfgx.Instance("normal")
		c.SetAdapter(normalAdpt)
		return c
	}()

	errorAdpt = &errorAdapter{}
	errorCfgx = func() *gcfgx.Config {
		c := gcfgx.Instance("error")
		c.SetAdapter(errorAdpt)
		return c
	}()

	testErr = gerrorx.ErrorString("error")
)

func Test_Instance(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		normalCfgx.SetErrorFn(nil)
		t.Assert(normalCfgx.MustGet(ctx, "key1", "value1"), "[key1]")
		t.Assert(normalCfgx.MustGetWithEnv(ctx, "key2", "value2"), "[key2]")
		t.Assert(normalCfgx.MustGetWithCmd(ctx, "key3", "value3"), "[key3]")
		t.Assert(normalCfgx.MustData(ctx)["key"], "value")

		errorCfgx.SetErrorFn(nil)
		t.Assert(errorCfgx.MustGet(ctx, "key1", "value1"), "value1")
		t.Assert(errorCfgx.MustGetWithEnv(ctx, "key2", "value2"), "value2")
		t.Assert(errorCfgx.MustGetWithCmd(ctx, "key3", "value3"), "value3")
		t.AssertNil(errorCfgx.MustData(ctx))
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

		errorCfgx.SetErrorFn(ffn)
		t.AssertNil(errorCfgx.MustGet(ctx, "key1", "value1"))
		t.AssertNil(errorCfgx.MustGetWithEnv(ctx, "key2", "value2"))
		t.AssertNil(errorCfgx.MustGetWithCmd(ctx, "key3", "value3"))
		t.AssertNil(errorCfgx.MustData(ctx))
	})
}

func Test_SetErrorLogger(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gcfgx.Instance().SetErrorLogger(g.Log()), nil)
		t.AssertNE(gcfgx.Instance().SetErrorLogger(nil), nil)
	})
}

type normalAdapter struct{}

func (a *normalAdapter) Available(_ context.Context, _ ...string) (ok bool) {
	return true
}

func (a *normalAdapter) Get(_ context.Context, pattern string) (value interface{}, err error) {
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
