package db

import (
	"calend/internal/modules/config"
	"calend/internal/modules/db/ent"
	"calend/internal/modules/db/trace_driver"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"go.uber.org/zap"
	"time"

	_ "calend/internal/modules/db/ent/runtime"
	// _ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func connectDB(cfg config.Config, logger *zap.Logger) (*ent.Client, error) {
	/* тут используем cgf и формируем строку подключения */
	connString := "host=localhost port=5432 user=postgres dbname=server_db password=passw0rd sslmode=disable"

	db, err := sql.Open(dialect.Postgres, connString)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к БД: %w", err)
	}

	logLevel := trace_driver.Warn
	if cfg.TraceSQLCommands {
		logLevel = trace_driver.Info
	}

	// Устанавливаем драйвер с трассировкой SQL команд
	traceDriver := trace_driver.NewTraceDriver(db, NewLogger(
		logger,
		trace_driver.Config{
			SlowThreshold: time.Duration(cfg.SQLSlowThreshold) * time.Second,
			LogLevel:      logLevel,
		}))

	// Формируем опции подключения
	var opts []ent.Option
	opts = append(opts, ent.Driver(traceDriver))

	client := ent.NewClient(opts...)

	return client, nil
}
