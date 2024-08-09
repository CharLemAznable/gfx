package gx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func GoIgnore(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
) {
	g.Go(ctx, goroutineFunc, nil)
}

func GoX(
	goroutineFunc func(),
	recoverFunc func(exception error),
) {
	if goroutineFunc == nil {
		return
	}
	go TryCatchX(goroutineFunc, recoverFunc)
}

func GoIgnoreX(
	goroutineFunc func(),
) {
	GoX(goroutineFunc, nil)
}

func TryIgnore(ctx context.Context, try func(ctx context.Context)) {
	_ = g.Try(ctx, try)
}

func TryX(try func()) (err error) {
	if try == nil {
		return
	}
	return g.Try(context.Background(), func(_ context.Context) { try() })
}

func TryCatchX(try func(), catch func(exception error)) {
	if try == nil {
		return
	}
	if exception := TryX(try); exception != nil && catch != nil {
		catch(exception)
	}
}

func TryIgnoreX(try func()) {
	_ = TryX(try)
}
