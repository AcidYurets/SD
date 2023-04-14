package http

import (
	"calend/internal/modules/config"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "HEAD", "OPTIONS"},
		Debug:            true,
	}).Handler)

	return r
}

func InvokeServer(r *mux.Router, config config.Config, logger *zap.Logger, lc fx.Lifecycle) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.HTTPServerHost, config.HTTPServerPort),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server")
			go func() {
				if err := server.ListenAndServe(); err != nil {
					logger.Sugar().Error(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server")
			return server.Shutdown(ctx)
		},
	})
}
