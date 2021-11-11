package logger

import (
	"context"

	grpc_zapctx "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
)

type ctxMarker struct{}

var (
	ctxMarkerKey = &ctxMarker{}
	nullLogger   = NewNopLogger()
)

func ToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ctxMarkerKey, logger)
}

func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(ctxMarkerKey).(Logger); ok {
		return logger
	}

	zlog := grpc_zapctx.Extract(ctx)
	sugar := zlog.Sugar()

	return &zapLogger{sugar}
}
