package gx_test

import (
	"github.com/CharLemAznable/gfx/frame/gx"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_Object(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gx.Client(), nil)
		t.AssertNE(gx.Server(), nil)
		t.AssertNE(gx.ViewX(), nil)
		t.AssertNE(gx.Config(), nil)
		t.AssertNE(gx.Cfg(), nil)
	})
}
