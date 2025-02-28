package gx

import "context"

func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
	finallyFunc ...func(ctx context.Context),
) {
	if goroutineFunc == nil {
		return
	}
	go TryCatch(ctx, goroutineFunc, recoverFunc, finallyFunc...)
}

func GoAnyway(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
) {
	Go(ctx, goroutineFunc, nil)
}

func GoX(
	goroutineFunc func(),
	recoverFunc func(exception error),
	finallyFunc ...func(),
) {
	if goroutineFunc == nil {
		return
	}
	go TryCatchX(goroutineFunc, recoverFunc, finallyFunc...)
}

func GoAnywayX(
	goroutineFunc func(),
) {
	GoX(goroutineFunc, nil)
}
