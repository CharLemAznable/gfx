package utils_test

import (
	"github.com/CharLemAznable/gfx/internal/utils"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_DefaultOrNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(utils.DefaultOrNil[any]())
		t.Assert(utils.DefaultOrNil[any]("test").String(), "test")
		t.Assert(utils.DefaultOrNil[any]("test", "demo").String(), "test")
	})
}
