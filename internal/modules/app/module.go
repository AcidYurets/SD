package app

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewAppInfo)
	Invokables = fx.Invoke()
)
