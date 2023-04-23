package logger

import (
	"context"
	"go.uber.org/zap"
)

type loggerCtx struct{}

func GetFromCtx(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(loggerCtx{}).(*zap.Logger)
	if !ok {
		panic("логгер отсутствует в контексте")
	}

	return logger
}

func SetToCtx(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerCtx{}, logger)
}
