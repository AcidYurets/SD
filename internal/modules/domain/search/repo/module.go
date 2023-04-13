package repo

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewSearchRepo)
	Invokables = fx.Invoke()
)
