package generator

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewTagGenerator)
	Invokables = fx.Invoke()
)
