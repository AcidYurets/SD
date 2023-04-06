package integration

import (
	"calend/internal/modules/app"
	"calend/internal/modules/config"
	"calend/internal/modules/db"
	"calend/internal/modules/domain"
	"calend/internal/modules/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var (
	TestModule = fx.Options(
		app.Module,
		logger.Module,
		config.Module,
		db.Module,
		domain.Module,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	TestInvokables = fx.Options(
		app.Invokables,
		logger.Invokables,
		config.Invokables,
		db.Invokables,
		domain.Invokables,
	)
)
