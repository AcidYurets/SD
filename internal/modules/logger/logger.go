package logger

import (
	"calend/internal/modules/app"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(app app.App) (*zap.Logger, zap.AtomicLevel, error) {
	loggerLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level = loggerLevel
	logger, err := config.Build()

	return logger, loggerLevel, err
}

func InvokeLogger(logger *zap.Logger, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return nil //logger.Sync()
		},
	})
}
