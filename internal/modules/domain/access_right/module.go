package access_right

import (
	"calend/internal/modules/domain/access_right/repo"
	"calend/internal/modules/domain/access_right/service"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.AccessRightRepo) *repo.AccessRightRepo { return r },
				fx.As(new(service.IAccessRightRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
