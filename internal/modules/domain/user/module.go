package user

import (
	"calend/internal/modules/domain/user/repo"
	"calend/internal/modules/domain/user/service"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.UserRepo) *repo.UserRepo { return r },
				fx.As(new(service.IUserRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)