package http

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewRouter)
	Invokables = fx.Invoke(InvokeServer)
)
