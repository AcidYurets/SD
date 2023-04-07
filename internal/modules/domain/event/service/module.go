package service

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewEventService)
	Invokables = fx.Invoke()
)
