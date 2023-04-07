package repo

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewEventRepo, NewInvitationRepo)
	Invokables = fx.Invoke()
)
