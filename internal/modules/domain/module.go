package domain

import (
	"calend/internal/modules/domain/access_right"
	"calend/internal/modules/domain/auth"
	"calend/internal/modules/domain/event"
	"calend/internal/modules/domain/tag"
	"calend/internal/modules/domain/user"
	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		access_right.Module,
		auth.Module,
		event.Module,
		tag.Module,
		user.Module,
	)
	Invokables = fx.Options(
		access_right.Invokables,
		auth.Invokables,
		event.Invokables,
		tag.Invokables,
		user.Invokables,
	)
)
