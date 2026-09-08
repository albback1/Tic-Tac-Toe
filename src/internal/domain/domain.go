package domain

import (
	"github.com/google/uuid"
)

// Тип данных обозначающий ячейку игрового поля
type Cell int

const (
	Empty Cell = iota //Пустая ячейка
	X                 // Значения Х
	O                 // Значение О
)

// Тип данных о состоянии игры
type GameResult int

const (
	Ongoing GameResult = iota // Игра продолжается
	XWin                      // Победа игрока Х
	OWin                      // Победа игрока У
	Draw                      // Ничья
)

// Игровое поле
type Field [][]Cell

// Модель игры
type Game struct {
	UUID  uuid.UUID
	Field Field
}

// Конструктор игры
func NewGame(sizefield int, uuid uuid.UUID) *Game {

	f := make([][]Cell, sizefield)
	for i := 0; i < sizefield; i++ {
		f[i] = make([]Cell, sizefield)
	}

	return &Game{
		UUID:  uuid, // инициализация uuid игры
		Field: f,
	}
}
