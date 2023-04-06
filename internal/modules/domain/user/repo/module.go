package repo

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewUserRepo)
	Invokables = fx.Invoke()
)
