package apollo_test

import (
	"context"
	"github.com/CharLemAznable/gfx/ext/gcfg/apollo"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_New_Default(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		adapter, err := apollo.LoadAdapter(ctx)
		t.AssertNE(adapter, nil)
		t.AssertNil(err)
	})
}

func Test_New_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		adapter, err := apollo.LoadAdapter(ctx, "apollo_error")
		t.AssertNil(adapter)
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "The Key field is required")
	})
}
