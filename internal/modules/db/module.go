package db

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewDBClient)
	Invokables = fx.Invoke(InvokeDBClient)
)
