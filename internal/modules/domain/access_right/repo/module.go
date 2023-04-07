package repo

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewAccessRightRepo)
	Invokables = fx.Invoke()
)
