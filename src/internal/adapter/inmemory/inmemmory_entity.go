package inmemory

import (
	"sync"

	"github.com/google/uuid"
)

// Игровое поле
type FieldRep [][]int

// Модель игры для хранения в inmemmory
type GameRep struct {
	Uuid  uuid.UUID
	Field FieldRep
}

// Структура хранилищя inmemmory
type InMemory struct {
	storage sync.Map
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

// конструктор модели GameRep
func NewGameRep(sizefield int, uuid uuid.UUID) *GameRep {
	f := make([][]int, sizefield)
	for i := 0; i < sizefield; i++ {
		f[i] = make([]int, sizefield)
	}

	return &GameRep{
		Uuid:  uuid,
		Field: f,
	}

}
