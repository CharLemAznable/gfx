package gcfgx_test

import (
	"github.com/CharLemAznable/gfx/os/gcfgx"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_AdapterDefault_Normal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		adapterContent, _ := gcfg.NewAdapterContent(`key: "value"`)
		adapterDefault := gcfgx.NewAdapterDefault(adapterContent)
		t.Assert(adapterDefault.Available(ctx), true)
		val, err := adapterDefault.Get(ctx, "key")
		t.Assert(val, "value")
		t.AssertNil(err)
		data, err := adapterDefault.Data(ctx)
		t.Assert(data, map[string]interface{}{"key": "value"})
		t.AssertNil(err)
	})
}

func Test_AdapterDefault_Fallback(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		adapterContent, _ := gcfg.NewAdapterContent()
		adapterFallback, _ := gcfg.NewAdapterContent(`key: "value"`)
		adapterDefault := gcfgx.NewAdapterDefault(adapterContent, adapterFallback)
		t.Assert(adapterDefault.Available(ctx), true)
		val, err := adapterDefault.Get(ctx, "key")
		t.Assert(val, "value")
		t.AssertNil(err)
		data, err := adapterDefault.Data(ctx)
		t.Assert(data, map[string]interface{}{"key": "value"})
		t.AssertNil(err)
	})
}

func Test_AdapterDefault_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		adapterContent, _ := gcfg.NewAdapterContent()
		adapterFallback, _ := gcfg.NewAdapterContent()
		adapterDefault := gcfgx.NewAdapterDefault(adapterContent, adapterFallback)
		t.Assert(adapterDefault.Available(ctx), false)
		val, err := adapterDefault.Get(ctx, "key")
		t.AssertNil(val)
		t.AssertNil(err)
		data, err := adapterDefault.Data(ctx)
		t.AssertNil(data)
		t.AssertNil(err)
	})
}
