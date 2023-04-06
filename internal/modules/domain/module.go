package domain

import (
	"calend/internal/modules/domain/auth"
	"calend/internal/modules/domain/user"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		user.Module,
		auth.Module,
	)
	Invokables = fx.Options(
		user.Invokables,
		auth.Invokables,
	)
)
