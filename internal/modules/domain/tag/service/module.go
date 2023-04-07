package service

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewTagService)
	Invokables = fx.Invoke()
)
