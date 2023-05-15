package generator

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewUserGenerator)
	Invokables = fx.Invoke()
)
