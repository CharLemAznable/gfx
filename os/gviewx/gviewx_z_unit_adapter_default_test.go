package gviewx_test

import (
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_AdapterDefault_Normal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testdata := gviewx.NewAdapterFile("testdata")
		adapterDefault := gviewx.NewAdapterDefault(testdata)
		content, err := adapterDefault.GetContent("test1")
		t.Assert(content, "Hello, {{.Name}}!")
		t.AssertNil(err)
	})
}

func Test_AdapterDefault_Fallback(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testdata := gviewx.NewAdapterFile()
		fallback := gviewx.NewAdapterFile("testdata")
		adapterDefault := gviewx.NewAdapterDefault(testdata, fallback)
		content, err := adapterDefault.GetContent("test1")
		t.Assert(content, "Hello, {{.Name}}!")
		t.AssertNil(err)
	})
}

func Test_AdapterDefault_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testdata := gviewx.NewAdapterFile()
		fallback := gviewx.NewAdapterFile()
		adapterDefault := gviewx.NewAdapterDefault(testdata, fallback)
		content, err := adapterDefault.GetContent("test1")
		t.Assert(content, "")
		t.Assert(gerror.Code(err), gcode.CodeInvalidParameter)
	})
}
