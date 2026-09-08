package dto

// запрос на ход (от игрока)
type MoveRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// ответ сервера
type GameResponse struct {
	ID       string    `json:"id"`
	Field    [3][3]int `json:"field"`
	Winner   int       `json:"winner"`
	GameOver bool      `json:"game_over"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
