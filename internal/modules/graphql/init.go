package graphql

import (
	"calend/internal/models/roles"
	"calend/internal/models/session"
	"calend/internal/modules/domain/auth/service"
	"calend/internal/modules/graphql/generated"
	"calend/internal/modules/graphql/resolvers"
	"calend/internal/modules/logger"
	"calend/internal/utils/slice"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/mux"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"net/http"
	"time"
)

//go:generate go run -mod=mod github.com/99designs/gqlgen@v0.17.29

func RegisterGraphQL(router *mux.Router, resolver *resolvers.Resolver, authService *service.AuthService, logger *zap.Logger) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		defaultErr := graphql.DefaultErrorPresenter(ctx, e)
		return defaultErr
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return fmt.Errorf("%v", err)
	})

	srv.AroundOperations(injectSession(
		// Операции, не требующие авторизации
		[]string{"SignUp", "Login", "RefreshToken", "IntrospectionQuery"},
		authService,
	))
	srv.AroundOperations(injectLogger(logger))
	// queryLogger последний, т.к. именно он выполняет саму операцию
	srv.AroundOperations(queryLogger(logger))

	// Раздаем содержимое папки static с gql песочницей
	router.PathPrefix("/sandbox").Handler(http.FileServer(http.Dir("./static")))
	router.Handle("/graphql", srv).Methods("GET", "POST", "OPTIONS")
}

// injectSession формирует функцию, которая проверяет токен и при успехе добавляет сессию в контекст
func injectSession(permittedOperations []string, authService *service.AuthService) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		opCtx := graphql.GetOperationContext(ctx)
		query := opCtx.OperationName

		if slice.Contains(permittedOperations, query) {
			// Устанавливаем в контекст необходимость использовать суперпользователю в Postgres
			ctx = roles.SetUseSuperUser(ctx)
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

		// Устанавливаем в контекст сессию
		ctx = session.SetSessionToCtx(ctx, *ss)
		// Устанавливаем в контекст необходимость смены пользователя перед каждым запросом
		ctx = roles.SetNeedChange(ctx)

		return next(ctx)
	}
}

// injectLogger формирует функцию, которая добавляет логгер в контекст
func injectLogger(lg *zap.Logger) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = logger.SetToCtx(ctx, lg)

		return next(ctx)
	}
}

// queryLogger формирует функцию, которая логгирует каждый входящий запрос
func queryLogger(lg *zap.Logger) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		var start, stop time.Time

		respHandler := next(ctx)
		oc := graphql.GetOperationContext(ctx)

		start = time.Now()
		resp := respHandler(ctx)
		stop = time.Now()

		fields := []zap.Field{
			zap.String("operation_name", oc.OperationName),
			zap.String("raw_query", oc.RawQuery),
			zap.Any("variables", oc.Variables),
			zap.String("latency", stop.Sub(start).String()),
		}

		if resp.Errors == nil {
			lg.Info("запрос выполнен успешно", fields...)

		} else {
			errs := make([]error, len(resp.Errors))
			for i, err := range resp.Errors {
				errs[i] = err
			}

			fields = append(fields, zap.Errors("ошибки", errs))
			lg.Error("запрос выполнен с ошибкой", fields...)
		}

		return func(_ context.Context) *graphql.Response {
			return resp
		}
	}
}
