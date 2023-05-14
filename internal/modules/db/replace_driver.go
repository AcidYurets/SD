package db

import (
	"calend/internal/modules/config"
	"calend/internal/modules/db/trace_driver"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
)

type DriverHolder interface {
	Driver() dialect.Driver
	SetDriver(dialect.Driver)
}

func replaceDriver(ctx context.Context, holder DriverHolder, cfg config.Config, logger trace_driver.ILogger) (*sql.Conn, error) {
	traceDriver, ok := holder.Driver().(*trace_driver.TraceDriver)
	if !ok {
		return nil, fmt.Errorf("некорректый тип драйвера - требуется *TraceDriver")
	}

	driver, ok := traceDriver.Driver.(*entsql.Driver)
	if !ok {
		return nil, fmt.Errorf("некорректый тип драйвера - требуется *entsql.Driver")
	}

	conn, err := driver.DB().Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить соединение")
	}
	entConn := entsql.Conn{ExecQuerier: conn}

	newDriver := entsql.NewDriver(cfg.DBDriver, entConn)
	newTraceDriver := trace_driver.NewTraceDriver(newDriver, logger)

	holder.SetDriver(newTraceDriver)

	return conn, nil
}
