package trace_driver

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const argsTypeError = "ошибка при логировании: args не является типом []any"

// LogLevel
type LogLevel int

const (
	Silent LogLevel = iota + 1
	Error
	Warn
	Info
)

type Config struct {
	SlowThreshold time.Duration // Длительность для определения долгих команд, используется при уровне Warn
	LogLevel      LogLevel      // Уровень логирования
}

// ILogger logger interface
type ILogger interface {
	LogMode(LogLevel) ILogger
	Error(msg string, data map[string]any)
	Warn(msg string, data map[string]any)
	Info(msg string, data map[string]any)
	Trace(begin time.Time, fc func() (sql string), err error, data map[string]any)
}

// TraceDriver is a driver that logs all driver operations.
type TraceDriver struct {
	dialect.Driver
	logger ILogger
}

// NewTraceDriver gets a driver and logger, and returns
// a new trace-driver that prints all outgoing operations.
func NewTraceDriver(d dialect.Driver, logger ILogger) dialect.Driver {
	drv := &TraceDriver{d, logger}
	return drv
}

func explainFunc(query string, args []any) func() (sql string) {
	return func() string {
		return Explain(query, args...)
	}
}

func errorFunc(msg string) func() string {
	return func() string {
		return msg
	}
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *TraceDriver) Exec(ctx context.Context, query string, args, v any) error {
	var f func() string
	argsSlice, ok := args.([]any)
	if !ok {
		f = errorFunc(argsTypeError)
	} else {
		f = explainFunc(query, argsSlice)
	}

	begin := time.Now()
	err := d.Driver.Exec(ctx, query, args, v)
	d.logger.Trace(begin, f, err, nil)

	return err
}

// ExecContext logs its params and calls the underlying driver ExecContext method if it is supported.
func (d *TraceDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("dialect.Driver.ExecContext is not supported")
	}

	f := explainFunc(query, args)
	begin := time.Now()
	res, err := drv.ExecContext(ctx, query, args...)
	d.logger.Trace(begin, f, err, nil)

	return res, err
}

// Query logs its params and calls the underlying driver Query method.
func (d *TraceDriver) Query(ctx context.Context, query string, args, v any) error {
	var f func() string
	argsSlice, ok := args.([]any)
	if !ok {
		f = errorFunc(argsTypeError)
	} else {
		f = explainFunc(query, argsSlice)
	}

	begin := time.Now()
	err := d.Driver.Query(ctx, query, args, v)
	d.logger.Trace(begin, f, err, nil)

	return err
}

// QueryContext logs its params and calls the underlying driver QueryContext method if it is supported.
func (d *TraceDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("dialect.Driver.QueryContext is not supported")
	}

	f := explainFunc(query, args)
	begin := time.Now()
	res, err := drv.QueryContext(ctx, query, args...)
	d.logger.Trace(begin, f, err, nil)

	return res, err
}

// Tx adds an log-id for the transaction and calls the underlying driver Tx command.
func (d *TraceDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	id := uuid.NewString()

	d.logger.Info("transaction started", map[string]any{"tx-id": id})

	return &DebugTx{tx, id, d.logger, ctx}, nil
}

// BeginTx adds an log-id for the transaction and calls the underlying driver BeginTx command if it is supported.
func (d *TraceDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	drv, ok := d.Driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	})
	if !ok {
		return nil, fmt.Errorf("dialect.Driver.BeginTx is not supported")
	}
	tx, err := drv.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()

	d.logger.Info("transaction started", map[string]any{"tx-id": id})

	return &DebugTx{tx, id, d.logger, ctx}, nil
}

// DebugTx is a transaction implementation that logs all transaction operations.
type DebugTx struct {
	dialect.Tx
	id     string
	logger ILogger
	ctx    context.Context
}

// Exec logs its params and calls the underlying transaction Exec method.
func (d *DebugTx) Exec(ctx context.Context, query string, args, v any) error {
	var f func() string
	argsSlice, ok := args.([]any)
	if !ok {
		f = errorFunc(argsTypeError)
	} else {
		f = explainFunc(query, argsSlice)
	}

	begin := time.Now()
	err := d.Tx.Exec(ctx, query, args, v)
	d.logger.Trace(begin, f, err, map[string]any{"tx-id": d.id})

	return err
}

// ExecContext logs its params and calls the underlying transaction ExecContext method if it is supported.
func (d *DebugTx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Tx.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.ExecContext is not supported")
	}

	f := explainFunc(query, args)
	begin := time.Now()
	res, err := drv.ExecContext(ctx, query, args...)
	d.logger.Trace(begin, f, err, map[string]any{"tx-id": d.id})

	return res, err
}

// Query logs its params and calls the underlying transaction Query method.
func (d *DebugTx) Query(ctx context.Context, query string, args, v any) error {
	var f func() string
	argsSlice, ok := args.([]any)
	if !ok {
		f = errorFunc(argsTypeError)
	} else {
		f = explainFunc(query, argsSlice)
	}

	begin := time.Now()
	err := d.Tx.Query(ctx, query, args, v)
	d.logger.Trace(begin, f, err, map[string]any{"tx-id": d.id})

	return err
}

// QueryContext logs its params and calls the underlying transaction QueryContext method if it is supported.
func (d *DebugTx) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Tx.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.QueryContext is not supported")
	}

	f := explainFunc(query, args)
	begin := time.Now()
	res, err := drv.QueryContext(ctx, query, args...)
	d.logger.Trace(begin, f, err, map[string]any{"tx-id": d.id})

	return res, err
}

// Commit logs this step and calls the underlying transaction Commit method.
func (d *DebugTx) Commit() error {
	d.logger.Info("transaction committed", map[string]any{"tx-id": d.id})

	return d.Tx.Commit()
}

// Rollback logs this step and calls the underlying transaction Rollback method.
func (d *DebugTx) Rollback() error {
	d.logger.Info("transaction rollbacked", map[string]any{"tx-id": d.id})

	return d.Tx.Rollback()
}
