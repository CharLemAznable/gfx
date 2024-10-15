package gx_test

import (
	"context"
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

func Test_Func(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gx.GoIgnore(context.Background(), func(ctx context.Context) {
			panic("ignored")
		})

		gx.GoX(nil, func(exception error) {
			panic("cannot happen")
		})
		gx.GoX(func() {
			panic("caught")
		}, func(exception error) {
			t.Assert(exception.Error(), "caught")
		})

		gx.GoIgnoreX(func() {
			panic("ignored")
		})

		gx.TryIgnore(context.Background(), func(ctx context.Context) {
			panic("ignored")
		})

		t.AssertNil(gx.TryX(nil))
		gx.TryCatchX(nil, func(exception error) {
			panic("cannot happen")
		})
		gx.TryCatchX(func() {
			panic("caught")
		}, func(exception error) {
			t.Assert(exception.Error(), "caught")
		})

		gx.TryIgnoreX(func() {
			panic("ignored")
		})
	})
}
