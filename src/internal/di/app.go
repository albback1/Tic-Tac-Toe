package di

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"src/internal/datasource/repository"
	"src/internal/domain/service"
	"src/internal/web/app"
	"src/internal/web/route"
)

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			zap.NewExample,
			repository.NewGameRepository,
			fx.Annotate(
				service.NewMinimaxService,
				fx.As(new(service.GameService)),
			),
			app.NewGameHandler,
		),
		fx.Invoke(RegisterHTTPServer),
	)
}

func RegisterHTTPServer(
	lifecycle fx.Lifecycle,
	handler *app.GameHandler,
	logger *zap.Logger,
) {
	route.SetupRoutes(handler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting HTTP server", zap.String("addr", server.Addr))
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatal("Failed to start server", zap.Error(err))
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
