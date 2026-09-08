package model

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	UUID      string    `json:"id"`
	Field     Field     `json:"field"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewGame() *Game {
	return &Game{
		UUID:      uuid.New().String(),
		Field:     *NewField(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewGameWithID(id string, field Field) *Game {
	now := time.Now()
	return &Game{
		UUID:      id,
		Field:     field,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (game *Game) GetWinner() int {
	return game.Field.GetWinner()
}

func (game *Game) CheckGameOver() bool {
	winner := game.GetWinner()
	if winner != 0 {
		return true
	}
	return game.Field.IsDraw()
}

func (game *Game) UpdateField(newField Field) {
	game.Field = newField
	game.UpdatedAt = time.Now()
}
