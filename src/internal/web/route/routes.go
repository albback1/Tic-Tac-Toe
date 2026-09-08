package route

import (
	"net/http"
	"src/internal/web/app"
)

func SetupRoutes(handler *app.GameHandler) {
	http.Handle("/game/", handler)
}
