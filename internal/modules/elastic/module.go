package elastic

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewClient)
	Invokables = fx.Invoke()
)
