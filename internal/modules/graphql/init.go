package graphql

import (
	"calend/internal/models/session"
	"calend/internal/modules/domain/auth/service"
	"calend/internal/modules/graphql/generated"
	"calend/internal/modules/graphql/resolvers"
	"calend/internal/utils/slice"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/mux"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

//go:generate go run -mod=mod github.com/99designs/gqlgen@v0.17.29

func RegisterGraphQL(router *mux.Router, resolver *resolvers.Resolver, authService *service.AuthService) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		defaultErr := graphql.DefaultErrorPresenter(ctx, e)
		return defaultErr
	})

	permittedOperations := []string{"SignUp", "Login", "RefreshToken", "IntrospectionQuery"}
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		opCtx := graphql.GetOperationContext(ctx)
		query := opCtx.OperationName

		if slice.Contains(permittedOperations, query) {
			// Разрешаем пройти дальне без авторизации
			return next(ctx)
		}

		// Получаем токен
		tkn := opCtx.Headers.Get("Authorization")
		if tkn == "" {
			return func(ctx context.Context) *graphql.Response {
				return graphql.ErrorResponse(ctx, "токен не указан")
			}
		}

		ss, err := authService.AuthAccessToken(ctx, tkn)
		if err != nil {
			return func(ctx context.Context) *graphql.Response {
				return graphql.ErrorResponse(ctx, err.Error())
			}
		}

		ctx = session.SetSessionToCtx(ctx, *ss)

		return next(ctx)
	})

	// Раздаем содержимое папки static с gql песочницей
	router.PathPrefix("/sandbox").Handler(http.FileServer(http.Dir("./static")))
	router.Handle("/graphql", srv).Methods("GET", "POST", "OPTIONS")
}
