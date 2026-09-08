package di

import (
	"net/http"

	"go.uber.org/fx"
)

type Route interface {
	http.Handler
	Pattern() string
}

// AsRoute аннотирует хендлер как Route для группы
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
