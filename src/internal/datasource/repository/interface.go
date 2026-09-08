package repository

// интерфейс сервиса в datasource (интерфейсс репозитория)

import "src/internal/domain/model"

type GameRepository interface {
	SaveGame(game *model.Game) error
	GetGame(id string) (*model.Game, error)
}
