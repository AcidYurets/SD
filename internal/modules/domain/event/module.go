package event

import (
	"calend/internal/modules/domain/event/repo"
	"calend/internal/modules/domain/event/service"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.EventRepo) *repo.EventRepo { return r },
				fx.As(new(service.IEventRepo)),
			),
			fx.Annotate(
				func(r *repo.InvitationRepo) *repo.InvitationRepo { return r },
				fx.As(new(service.IInvitationRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
