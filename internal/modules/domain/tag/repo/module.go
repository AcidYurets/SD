package repo

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewTagRepo)
	Invokables = fx.Invoke()
)
