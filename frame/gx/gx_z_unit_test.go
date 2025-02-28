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

func Test_Goroutine(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gx.Go(context.Background(), nil, func(ctx context.Context, exception error) {
			panic("cannot happen")
		})
		gx.GoAnyway(context.Background(), func(ctx context.Context) {
			panic("ignored")
		})

		gx.GoX(nil, func(exception error) {
			panic("cannot happen")
		})
		gx.GoAnywayX(func() {
			panic("ignored")
		})

		gx.GoX(func() {
			panic("caught")
		}, func(exception error) {
			t.Assert(exception.Error(), "caught")
		})
	})
}

func Test_Try_Catch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gx.TryIgnore(context.Background(), func(ctx context.Context) {
			panic("ignored")
		})

		t0 := 0
		t.AssertNil(gx.TryX(nil, func() { t0 = 1 }))
		t.Assert(gx.TryX(func() { panic("caught") }, func() { t0 = 2 }).Error(), "caught")
		t.Assert(t0, 2)

		gx.TryIgnoreX(func() {
			panic("ignored")
		})

		t1 := 0
		gx.TryCatchX(nil, func(exception error) {
			t1 = 1
			panic("cannot happen")
		}, func() {
			t1 = 2
		})
		gx.TryCatchX(func() {
			panic("caught")
		}, func(exception error) {
			t1 = 3
			t.Assert(exception.Error(), "caught")
		}, func() {
			t1 = 4
		})
		t.Assert(t1, 4)
	})
}
