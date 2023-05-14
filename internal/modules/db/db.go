package db

import (
	"calend/internal/modules/config"
	"calend/internal/modules/db/ent"
	_ "calend/internal/modules/db/ent/runtime"
	"calend/internal/modules/db/roles"
	"calend/internal/modules/db/trace_driver"
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent --feature  sql/execquery,intercept,schema/snapshot --template ./templates ./schema

func NewDBClient(cfg config.Config, logger *zap.Logger) (*ent.Client, trace_driver.ILogger, error) {
	client, traceLogger, err := connectDB(cfg, logger)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при подключении к базе данных: %w", err)
	}

	return client, traceLogger, nil
}

func InvokeDBClient(
	client *ent.Client,
	cfg config.Config,
	logger trace_driver.ILogger,
	lifecycle fx.Lifecycle,
) error {
	// Делаем миграцию при необходимости
	if cfg.AutoMigrate {
		if err := client.Schema.Create(context.Background()); err != nil {
			return fmt.Errorf("ошибка при миграции: %w", err)
		}
	}

	_, err := client.ExecContext(context.Background(), migrationQuery)
	if err != nil {
		return err
	}

	interceptor := ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			drvHolder, ok := query.(DriverHolder)
			if !ok {
				return nil, fmt.Errorf("query не реализует методы Driver или SetDriver")
			}

			conn, err := replaceDriver(ctx, drvHolder, cfg, logger)
			if err != nil {
				return nil, err
			}

			err = roles.ChangeRole(ctx, conn)
			if err != nil {
				return nil, err
			}

			value, err := next.Query(ctx, query)
			if err != nil {
				return nil, err
			}

			err = conn.Close()
			if err != nil {
				return nil, err
			}

			return value, err
		})
	})

	hook := func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			drvHolder, ok := mutation.(DriverHolder)
			if !ok {
				return nil, fmt.Errorf("mutation не реализует методы Driver или SetDriver")
			}

			conn, err := replaceDriver(ctx, drvHolder, cfg, logger)
			if err != nil {
				return nil, err
			}

			err = roles.ChangeRole(ctx, conn)
			if err != nil {
				return nil, err
			}

			value, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			err = conn.Close()
			if err != nil {
				return nil, err
			}

			return value, err
		})
	}

	client.Intercept(interceptor)
	client.Use(hook)

	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return client.Close()
		},
	})

	return nil
}
