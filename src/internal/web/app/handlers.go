package app

// хэндлер с net/http

import (
	"encoding/json"
	"net/http"
	"strings"

	"src/internal/domain/service"
	"src/internal/web/dto"
	"src/internal/web/mapper"
)

type GameHandler struct {
	gameService service.GameService
}

func NewGameHandler(gameService service.GameService) *GameHandler {
	return &GameHandler{gameService: gameService}
}

// Pattern возвращает путь для маршрутизации
func (handler *GameHandler) Pattern() string {
	return "/game/"
}

// Implement http.Handler interface
func (handler *GameHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Extract game ID from path
	path := strings.TrimPrefix(request.URL.Path, "/game/")
	gameID := strings.SplitN(path, "/", 2)[0]

	switch request.Method {
	case http.MethodPost:
		if gameID == "" {
			game, err := handler.gameService.CreateGame()
			if err != nil {
				handler.RespondWithError(writer, err.Error(), 500)
				return
			}
			handler.RespondWithJSON(writer, mapper.ToResponse(game), 200)
			return
		}
		handler.PostGame(writer, request, gameID)
	default:
		handler.RespondWithError(writer, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *GameHandler) PostGame(writer http.ResponseWriter, request *http.Request, gameID string) {
	// читаем JSON
	var req dto.MoveRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		handler.RespondWithError(writer, "invalid JSON", http.StatusBadRequest)
		return
	}

	// проверяем игру
	err := handler.gameService.ValidateFieldGame(gameID, req.Row, req.Col)
	if err != nil {
		handler.RespondWithError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// обработка хода (игрок + ИИ)
	updatedGame, err := handler.gameService.MakeMove(gameID, req.Row, req.Col)
	if err != nil {
		handler.RespondWithError(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// ответ
	response := mapper.ToResponse(updatedGame)
	handler.RespondWithJSON(writer, response, http.StatusOK)
}

func (handler *GameHandler) RespondWithError(writer http.ResponseWriter, message string, status int) {
	handler.RespondWithJSON(writer, dto.ErrorResponse{Error: message}, status)
}

func (handler *GameHandler) RespondWithJSON(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(data)
}
