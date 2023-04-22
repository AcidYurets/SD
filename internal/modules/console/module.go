package console

import "go.uber.org/fx"

var (
	Module     = fx.Provide()
	Invokables = fx.Invoke(MenuLoop)
)
