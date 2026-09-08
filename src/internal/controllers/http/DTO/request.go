package dto

import (
	"github.com/google/uuid"
)

// стурктура http запроса
type RequestGame struct {
	Uuid  uuid.UUID `json:"uuid,omitempty"`
	Field [][]int   `json:"field"`
}

func NewRequestGame(sizefield int, uuid uuid.UUID) *RequestGame {
	f := make([][]int, sizefield)
	for i := 0; i < sizefield; i++ {
		f[i] = make([]int, sizefield)
	}

	return &RequestGame{
		Uuid:  uuid, // инициализация uuid игры
		Field: f,
	}
}
