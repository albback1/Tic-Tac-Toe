package di

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"tictac/internal/adapter/inmemory"
	httpController "tictac/internal/controllers/http"
	"tictac/internal/domain"
	"tictac/internal/usecase"

	"go.uber.org/fx"
)

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				inmemory.NewInMemory,
				fx.As(new(usecase.GameRepository)),
			),
			domain.NewGame,
			usecase.NewGameServiceItml,
			AsRoute(httpController.NewGameHandler),
			fx.Annotate(
				NewServMux,
				fx.ParamTags(`group:"routes"`),
			),
			NewHttpServer,
		),
		fx.Invoke(func(*http.Server) {}),
	)
}

func NewHttpServer(ls fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	ls.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		}, OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func NewServMux(routes []httpController.Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Patern(), route)
	}
	return mux
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(httpController.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
