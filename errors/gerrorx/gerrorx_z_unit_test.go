package gerrorx_test

import (
	"github.com/CharLemAznable/gfx/errors/gerrorx"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_ErrorString(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gerrorx.ErrorString("test"), "test")
		t.Assert(gerrorx.ErrorString("test").Error(), "test")
	})
}
