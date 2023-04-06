package logger

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewLogger)
	Invokables = fx.Invoke(InvokeLogger)
)
