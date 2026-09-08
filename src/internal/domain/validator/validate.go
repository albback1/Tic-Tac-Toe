package validator

// валидация хода игрока

import (
	"errors"
	"src/internal/domain/model"
)

type GameValidator struct{}

func NewGameValidator() *GameValidator {
	return &GameValidator{}
}

func (validator *GameValidator) ValidateMove(field model.Field, row, column int) error {

	if row < 0 || row > 2 || column < 0 || column > 2 {
		return errors.New("invalid move position")
	}

	if field.Cells[row][column] != 0 {
		return errors.New("you cannot change this cell bc it is already occupied")
	}
	return nil

}

func (validator *GameValidator) ValidateGameOver(field model.Field) error {
	if field.GetWinner() != 0 || field.IsDraw() {
		return errors.New("game is over already")
	}
	return nil
}
