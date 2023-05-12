package elastic

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewEventElasticService)
	Invokables = fx.Invoke()
)
