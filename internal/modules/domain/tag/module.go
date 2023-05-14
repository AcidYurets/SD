package tag

import (
	"calend/internal/modules/domain/tag/generator"
	"calend/internal/modules/domain/tag/repo"
	"calend/internal/modules/domain/tag/service"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,
		generator.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.TagRepo) *repo.TagRepo { return r },
				fx.As(new(service.ITagRepo)),
				fx.As(new(generator.ITagRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
		generator.Invokables,
	)
)
