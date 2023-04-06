package auth

import (
	"calend/internal/modules/domain/auth/service"
	"calend/internal/modules/domain/user/repo"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.UserRepo) *repo.UserRepo { return r },
				fx.As(new(service.IUserRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
	)
)
