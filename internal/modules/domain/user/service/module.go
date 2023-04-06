package service

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewUserService)
	Invokables = fx.Invoke()
)
