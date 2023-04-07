package service

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewAccessRightService)
	Invokables = fx.Invoke()
)
