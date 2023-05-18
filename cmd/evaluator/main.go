package main

import (
	"calend/internal/modules"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		modules.EvaluatorModule,
		modules.EvaluatorInvokables,
	).Run()
}
