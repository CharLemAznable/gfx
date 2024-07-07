package gvarx_test

import (
	"github.com/CharLemAznable/gfx/container/gvarx"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_DefaultOrNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(gvarx.DefaultOrNil())
		t.Assert(gvarx.DefaultOrNil("test").String(), "test")
		t.Assert(gvarx.DefaultOrNil("test", "demo").String(), "test")
	})
}
