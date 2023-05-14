package main

import (
	"calend/internal/modules"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		modules.ConsoleAppModule,
		modules.ConsoleAppInvokables,
	).Run()
}
