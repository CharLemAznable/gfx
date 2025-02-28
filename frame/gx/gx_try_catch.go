package gx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func Try(ctx context.Context, try func(ctx context.Context), finally ...func(ctx context.Context)) (err error) {
	if len(finally) > 0 && finally[0] != nil {
		defer finally[0](ctx)
	}
	return g.Try(ctx, try)
}

func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error), finally ...func(ctx context.Context)) {
	if len(finally) > 0 && finally[0] != nil {
		defer finally[0](ctx)
	}
	g.TryCatch(ctx, try, catch)
}

func TryIgnore(ctx context.Context, try func(ctx context.Context), finally ...func(ctx context.Context)) {
	_ = Try(ctx, try, finally...)
}

func TryX(try func(), finally ...func()) (err error) {
	if try == nil {
		return
	}
	return Try(context.Background(), func(_ context.Context) { try() }, finallyFunc(finally...))
}

func TryCatchX(try func(), catch func(exception error), finally ...func()) {
	if try == nil {
		return
	}
	TryCatch(context.Background(), func(_ context.Context) { try() }, catchFunc(catch), finallyFunc(finally...))
}

func TryIgnoreX(try func(), finally ...func()) {
	_ = TryX(try, finally...)
}

func catchFunc(catch func(exception error)) (f func(ctx context.Context, exception error)) {
	if catch != nil {
		f = func(_ context.Context, exception error) { catch(exception) }
	}
	return
}

func finallyFunc(finally ...func()) (f func(ctx context.Context)) {
	if len(finally) > 0 && finally[0] != nil {
		f = func(_ context.Context) { finally[0]() }
	}
	return
}
