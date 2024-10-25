package utils_test

import (
	"github.com/CharLemAznable/gfx/internal/utils"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_DefaultOrNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(utils.DefaultOrNil())
		t.Assert(utils.DefaultOrNil("test").String(), "test")
		t.Assert(utils.DefaultOrNil("test", "demo").String(), "test")
	})
}

func Test_FormatString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.FormatCmdKey("TEST_DEMO"), "test.demo")
		t.Assert(utils.FormatEnvKey("test.demo"), "TEST_DEMO")
	})
}
