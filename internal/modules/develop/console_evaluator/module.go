package console_evaluator

import "go.uber.org/fx"

var (
	Module     = fx.Provide()
	Invokables = fx.Invoke(Eval)
)
