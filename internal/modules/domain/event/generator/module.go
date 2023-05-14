package generator

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewEventGenerator)
	Invokables = fx.Invoke()
)
