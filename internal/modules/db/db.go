package db

import (
	"calend/internal/modules/config"
	"calend/internal/modules/db/ent"
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// //go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent --template ./templates ./schema
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent --feature intercept,schema/snapshot ./schema

func NewDBClient(cfg config.Config, logger *zap.Logger) (*ent.Client, error) {
	client, err := connectDB(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к базе данных: %w", err)
	}

	return client, nil
}

func InvokeDBClient(
	client *ent.Client,
	cfg config.Config,
	lifecycle fx.Lifecycle,
) error {
	// Делаем миграцию при необходимости
	if cfg.AutoMigrate {
		if err := client.Schema.Create(context.Background()); err != nil {
			return fmt.Errorf("ошибка при миграции: %w", err)
		}
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return client.Close()
		},
	})

	return nil
}
