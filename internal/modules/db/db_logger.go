package db

import (
	"calend/internal/modules/db/ent"
	"calend/internal/modules/db/trace_driver"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func NewLogger(lg *zap.Logger, config trace_driver.Config) trace_driver.ILogger {
	return &logger{
		lg:     lg,
		Config: config,
	}
}

type logger struct {
	lg *zap.Logger
	trace_driver.Config
}

func (l *logger) LogMode(level trace_driver.LogLevel) trace_driver.ILogger {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l logger) Error(msg string, data map[string]any) {
	if l.LogLevel >= trace_driver.Error {
		l.lg.Error(msg, l.makeFields(data)...)
	}
}

func (l logger) Warn(msg string, data map[string]any) {
	if l.LogLevel >= trace_driver.Warn {
		l.lg.Warn(msg, l.makeFields(data)...)
	}
}
func (l logger) Info(msg string, data map[string]any) {
	if l.LogLevel >= trace_driver.Info {
		l.lg.Info(msg, l.makeFields(data)...)
	}
}

func (l logger) Trace(begin time.Time, fc func() string, err error, data map[string]any) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		strElapced := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)

		fields := l.makeFields(data)

		switch {
		case err != nil && l.LogLevel >= trace_driver.Error:
			var loggerFunc func(msg string, fields ...zap.Field)
			logger := l.lg.With(fields...)
			if errors.Is(err, &ent.NotFoundError{}) {
				loggerFunc = logger.Debug
			} else {
				loggerFunc = logger.Error
			}

			sql := fc()
			loggerFunc("Trace SQL",
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced),
				zap.Error(err))
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= trace_driver.Warn:
			sql := fc()
			l.lg.With(fields...).With(
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced),
				zap.Duration("slowThreshold", l.SlowThreshold)).
				Warn("Trace SQL")
		case l.LogLevel >= trace_driver.Info:
			sql := fc()
			l.lg.With(fields...).With(
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced)).
				Info("Trace SQL")
		}
	}
}

func (l logger) makeFields(data map[string]any) []zap.Field {
	var fields []zap.Field

	for k, v := range data {
		fields = append(fields, zap.Any(k, v))
	}

	return fields
}
