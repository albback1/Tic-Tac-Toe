package httpController

import (
	"encoding/json"
	"net/http"
	"strings"
	dto "tictac/internal/controllers/http/DTO"
	"tictac/internal/domain"
	"tictac/internal/usecase"

	"github.com/google/uuid"
)

func getGame(w http.ResponseWriter, servise *usecase.GameServiceItml) {
	uid := uuid.New()
	game := domain.NewGame(3, uid)
	servise.Repo.SaveGame(game)
	responce := dto.MapToDTO(game)
	responce.ResultGame = convertResultTypeToString(servise.GameOverCheck(game))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responce)
}

func postGame(w http.ResponseWriter, r *http.Request, servise *usecase.GameServiceItml) {
	defer r.Body.Close()

	uidStr := strings.TrimPrefix(r.URL.Path, "/game/")
	uid, err := uuid.Parse(uidStr) // парсем uuid игры
	if err != nil {                //проверяем валидный ли uuid
		http.Error(w, "Invalid uuid", http.StatusBadRequest)
		return
	}
	var request dto.RequestGame                    // заводим переменную dto реквест
	request.Uuid = uid                             // присваиваем uuid
	err = json.NewDecoder(r.Body).Decode(&request) //парсим тело запроса в DTO реквест
	if err != nil {                                // проверяем спарсился ли json
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newGame := dto.MapToDomain(&request) // мапим dto в domain

	bol, err := servise.ValidateField(newGame)

	if !bol && err == nil { //проверяем не изменляли ли поле
		http.Error(w, "invalid field", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	resultGame := servise.GameOverCheck(newGame)

	if resultGame == domain.Ongoing {
		x, y := servise.NextMoveMinimax(newGame, domain.X)          // определяем координаты следующего хода
		newGame.Field[y][x] = domain.X                              // делаем ход
		resultGame = servise.GameOverCheck(newGame)                 // проверяем итог игры
		servise.Repo.SaveGame(newGame)                              // сохраняем игру
		response := dto.MapToDTO(newGame)                           //мапим domain в dto
		response.ResultGame = convertResultTypeToString(resultGame) // сохраняем результат игры в DTo
		err = putResponseGame(w, response)                          //парсим dto в jso
		if err != nil {                                             //проверяем спарсилось ли в json
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		servise.Repo.SaveGame(newGame)                              //сохраняем игру
		response := dto.MapToDTO(newGame)                           //мапим domain в dto
		response.ResultGame = convertResultTypeToString(resultGame) // сохраняем результат игры в DTo
		err = putResponseGame(w, response)                          //парсим dto в jso
		if err != nil {                                             //проверяем спарсилось ли в json
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func putResponseGame(w http.ResponseWriter, response *dto.ResponseGame) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	return err
}

func convertResultTypeToString(resDomain domain.GameResult) string {

	switch resDomain {
	case domain.Draw:
		return "draw"
	case domain.OWin:
		return "player win"
	case domain.XWin:
		return "player lose"
	default:
		return "ongoing"
	}
}
