package repo_elastic

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewSearchRepo)
	Invokables = fx.Invoke()
)
