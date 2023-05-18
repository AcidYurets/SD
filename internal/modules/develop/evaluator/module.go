package evaluator

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewEvaluator)
	Invokables = fx.Invoke()
)
