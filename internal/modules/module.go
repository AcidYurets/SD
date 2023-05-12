package modules

import (
	"calend/internal/modules/app"
	"calend/internal/modules/config"
	"calend/internal/modules/console"
	"calend/internal/modules/db"
	"calend/internal/modules/domain"
	"calend/internal/modules/elastic"
	"calend/internal/modules/graphql"
	"calend/internal/modules/http"
	"calend/internal/modules/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var (
	AppModule = fx.Options(
		app.Module,
		logger.Module,
		config.Module,
		db.Module,
		elastic.Module,
		http.Module,
		graphql.Module,

		domain.Module,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	AppInvokables = fx.Options(
		app.Invokables,
		logger.Invokables,
		config.Invokables,
		db.Invokables,
		elastic.Invokables,
		http.Invokables,
		graphql.Invokables,

		domain.Invokables,
	)
)

var (
	ConsoleAppModule = fx.Options(
		app.Module,
		logger.Module,
		config.Module,
		db.Module,
		console.Module,

		domain.Module,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	ConsoleAppInvokables = fx.Options(
		app.Invokables,
		logger.Invokables,
		config.Invokables,
		db.Invokables,
		console.Invokables,

		domain.Invokables,
	)
)
