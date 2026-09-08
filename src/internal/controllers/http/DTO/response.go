package dto

import "github.com/google/uuid"

// структура http ответа
type ResponseGame struct {
	Uuid       uuid.UUID `json:"uuid"`
	Field      [][]int   `json:"field"`
	ResultGame string    `json:"resultgame"`
}

func NewResponseGame(sizefield int, uuid uuid.UUID) *ResponseGame {
	f := make([][]int, sizefield)
	for i := 0; i < sizefield; i++ {
		f[i] = make([]int, sizefield)
	}

	return &ResponseGame{
		Uuid:  uuid, // инициализация uuid игры
		Field: f,
	}
}
