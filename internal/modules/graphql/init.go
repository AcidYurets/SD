package graphql

import (
	"calend/internal/modules/graphql/generated"
	"calend/internal/modules/graphql/resolvers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/mux"
	"net/http"
)

//go:generate go run -mod=mod github.com/99designs/gqlgen@v0.17.29

func RegisterGraphQL(router *mux.Router, resolver *resolvers.Resolver) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	// Раздаем содержимое папки static с gql песочницей
	router.PathPrefix("/sandbox").Handler(http.FileServer(http.Dir("./static")))

	router.Handle("/graphql", srv).Methods("GET", "POST", "OPTIONS")
}
