package event

import (
	repo4 "calend/internal/modules/domain/access_right/repo"
	"calend/internal/modules/domain/event/elastic"
	"calend/internal/modules/domain/event/generator"
	"calend/internal/modules/domain/event/repo"
	"calend/internal/modules/domain/event/service"
	repo2 "calend/internal/modules/domain/tag/repo"
	repo3 "calend/internal/modules/domain/user/repo"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,
		elastic.Module,
		generator.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.EventRepo) *repo.EventRepo { return r },
				fx.As(new(service.IEventRepo)),
				fx.As(new(elastic.IEventRepo)),
				fx.As(new(generator.IEventRepo)),
			),
			fx.Annotate(
				func(r *repo.InvitationRepo) *repo.InvitationRepo { return r },
				fx.As(new(service.IInvitationRepo)),
				fx.As(new(generator.IInvitationRepo)),
			),

			fx.Annotate(
				func(r *repo2.TagRepo) *repo2.TagRepo { return r },
				fx.As(new(generator.ITagsRepo)),
			),

			fx.Annotate(
				func(r *repo3.UserRepo) *repo3.UserRepo { return r },
				fx.As(new(generator.IUserRepo)),
			),

			fx.Annotate(
				func(r *repo4.AccessRightRepo) *repo4.AccessRightRepo { return r },
				fx.As(new(generator.IAccessRightRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
		elastic.Invokables,
		generator.Invokables,
	)
)
