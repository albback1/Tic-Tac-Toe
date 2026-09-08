package service

// интерфейс сервиса в domain

import "src/internal/domain/model"

type GameService interface {
	MakeMove(gameID string, row, col int) (*model.Game, error) // получаем оьновленное поле и ход ИИ
	ValidateFieldGame(gameID string, row, column int) error    // проверяем, изменялись ли предыдущие ходы
	CheckGameOver(field model.Field) bool                      // ya know
	CreateGame() (*model.Game, error)
}
