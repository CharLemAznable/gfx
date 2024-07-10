package apollo_test

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/gcfg/apollo"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

var (
	ctx = context.TODO()
)

func Test_New_Default(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		adapter, err := apollo.LoadAdapter(ctx)
		t.AssertNE(adapter, nil)
		t.AssertNil(err)
	})
}

func Test_New_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		adapter, err := apollo.LoadAdapter(ctx, "apollo_error")
		t.AssertNil(adapter)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "The Key field is required")
	})
}

func Test_New_None(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("GF_GCFG_APOLLO_APPID", "APPID")
		defer func() { _ = genv.Remove("GF_GCFG_APOLLO_APPID") }()
		gcmd.Init([]string{"--gf.gcfg.apollo.ip=IP"}...)

		adapter, err := apollo.LoadAdapter(ctx, "apollo_none")
		t.AssertNil(adapter)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "The Key field is required")
	})
}
