package user

import (
	"calend/internal/modules/domain/user/service"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,
	)

	Invokables = fx.Options(
		service.Invokables,
	)
)
