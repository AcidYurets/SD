package elastic

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

func newLogger(lg *zap.Logger) elastic.Logger {
	return &logger{
		lg: lg,
	}
}

type logger struct {
	lg *zap.Logger
}

func (l *logger) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.lg.Info("Trace Elastic", zap.String("query", msg))
}
