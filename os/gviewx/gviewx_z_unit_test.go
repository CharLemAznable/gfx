package gviewx_test

import (
	"context"
	"github.com/CharLemAznable/gfx/os/gviewx"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

var (
	ctx = context.TODO()
)

func Test_Default(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		view := gviewx.Instance().
			SetAdapter(gviewx.NewAdapterFile("testdata"))
		result, err := view.Parse(ctx, "test1", g.Map{"Name": "John"})
		t.AssertNil(err)
		t.Assert(result, "Hello, John!")
		result, err = view.Parse(ctx, "test", g.Map{"Name": "John"})
		t.Assert(gerror.Code(err), gcode.CodeInvalidParameter)

		adapter := view.GetAdapter()
		content, err := adapter.GetContent("test1")
		t.AssertNil(err)
		t.Assert(content, "Hello, {{.Name}}!")
		content, err = adapter.GetContent("test")
		t.Assert(gerror.Code(err), gcode.CodeInvalidParameter)
	})
}

func Test_Config(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testAdapter, _ := gcfg.NewAdapterFile()
		_ = testAdapter.AddPath("testdata")
		oriAdapter := g.Config().GetAdapter()
		g.Config().SetAdapter(testAdapter)

		view := gviewx.Instance("test").
			SetAdapter(gviewx.NewAdapterFile("testdata")).
			SetConfig(gviewx.Config{
				I18nManager: gi18n.Instance(),
			})
		result, err := view.Parse(ctx, "test2", g.Map{"Name": "John"})
		t.AssertNil(err)
		t.Assert(result, "Hello again, Joe!")

		g.Config().SetAdapter(oriAdapter)
	})
}
