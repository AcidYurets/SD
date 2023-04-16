package graphql

import (
	"calend/internal/modules/graphql/resolvers"
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(resolvers.NewResolver)
	Invokables = fx.Invoke(RegisterGraphQL)
)
