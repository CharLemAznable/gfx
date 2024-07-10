package gvarx_test

import (
	"github.com/CharLemAznable/gfx/container/gvarx"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_DefaultOrNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(gvarx.DefaultOrNil())
		t.Assert(gvarx.DefaultOrNil("test").Val(), "test")
		t.Assert(gvarx.DefaultOrNil("test", "demo").Val(), "test")
	})
}

func Test_DefaultIfNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(gvarx.DefaultIfNil(nil))
		t.AssertNil(gvarx.DefaultIfNil(gvar.New(nil)))
		t.Assert(gvarx.DefaultIfNil(gvar.New("")).Val(), "")
		t.Assert(gvarx.DefaultIfNil(gvar.New(nil), "test").Val(), "test")
		t.Assert(gvarx.DefaultIfNil(gvar.New(""), "test").Val(), "")
		t.Assert(gvarx.DefaultIfNil(gvar.New("demo"), "test").Val(), "demo")
	})
}

func Test_DefaultIfEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(gvarx.DefaultIfEmpty(nil))
		t.AssertNil(gvarx.DefaultIfEmpty(gvar.New(nil)))
		t.AssertNil(gvarx.DefaultIfEmpty(gvar.New("")))
		t.Assert(gvarx.DefaultIfEmpty(gvar.New(nil), "test").Val(), "test")
		t.Assert(gvarx.DefaultIfEmpty(gvar.New(""), "test").Val(), "test")
		t.Assert(gvarx.DefaultIfEmpty(gvar.New("demo"), "test").Val(), "demo")
	})
}
