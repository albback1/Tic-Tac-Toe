package httpController

import (
	"net/http"
	"tictac/internal/usecase"
)

type Route interface {
	http.Handler

	Patern() string
}

type GameHandler struct {
	service *usecase.GameServiceItml
}

func NewGameHandler(service *usecase.GameServiceItml) *GameHandler {
	return &GameHandler{service: service}
}

func (h *GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getGame(w, h.service)
	case http.MethodPost:
		postGame(w, r, h.service)
	}
}

func (h *GameHandler) Patern() string {
	return "/game/"
}
