package search

import (
	"calend/internal/modules/domain/search/repo"
	"calend/internal/modules/domain/search/service"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.SearchRepo) *repo.SearchRepo { return r },
				fx.As(new(service.ISearchRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
