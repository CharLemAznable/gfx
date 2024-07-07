package gcmdx_test

import (
	"github.com/CharLemAznable/gfx/os/gcmdx"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_GetOpt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gcmd.Init([]string{"--key=value", "--test="}...)

		t.Assert(gcmd.GetOpt("key"), "value")
		t.Assert(gcmd.GetOpt("test"), "")
		t.Assert(gcmd.GetOpt("test", "x"), "")

		t.Assert(gcmdx.GetOpt("key"), "value")
		t.Assert(gcmdx.GetOpt("test"), "")
		t.Assert(gcmdx.GetOpt("test", "x"), "x")
	})
}

func Test_GetOptWithEnv(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("KEY", "value")
		_ = genv.Set("TEST", "")
		defer func() { _ = genv.Remove("KEY", "TEST") }()

		t.Assert(gcmd.GetOptWithEnv("key"), "value")
		t.Assert(gcmd.GetOptWithEnv("test"), "")
		t.Assert(gcmd.GetOptWithEnv("test", "x"), "")

		t.Assert(gcmdx.GetOptWithEnv("key"), "value")
		t.Assert(gcmdx.GetOptWithEnv("test"), "")
		t.Assert(gcmdx.GetOptWithEnv("test", "x"), "x")
	})
	gtest.C(t, func(t *gtest.T) {
		_ = genv.Set("TEST1", "value1")
		_ = genv.Set("TEST2", "")
		defer func() { _ = genv.Remove("TEST1", "TEST2") }()
		gcmd.Init([]string{"--key=value", "--test1="}...)

		t.Assert(gcmd.GetOptWithEnv("key"), "value")
		t.Assert(gcmd.GetOptWithEnv("test1"), "")
		t.Assert(gcmd.GetOptWithEnv("test1", "x"), "")
		t.Assert(gcmd.GetOptWithEnv("test2"), "")
		t.Assert(gcmd.GetOptWithEnv("test2", "y"), "")

		t.Assert(gcmdx.GetOptWithEnv("key"), "value")
		t.Assert(gcmdx.GetOptWithEnv("test1"), "value1")
		t.Assert(gcmdx.GetOptWithEnv("test1", "x"), "value1")
		t.Assert(gcmdx.GetOptWithEnv("test2"), "")
		t.Assert(gcmdx.GetOptWithEnv("test2", "y"), "y")
	})
}
