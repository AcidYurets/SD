package logger

import (
	"calend/internal/modules/app"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(app app.App) (*zap.Logger, zap.AtomicLevel, error) {
	logger, err := zap.NewProduction()
	loggerLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	return logger, loggerLevel, err
}

func InvokeLogger(logger *zap.Logger, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return logger.Sync()
		},
	})
}
