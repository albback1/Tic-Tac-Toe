package repository

// реализация интерфейса репозитория, потокобезопасно

import (
	"errors"
	"sync"

	dto "src/internal/datasource/dto"
	"src/internal/datasource/mapper"
	"src/internal/domain/model"
)

type gameRepository struct {
	storage *sync.Map
}

func NewGameRepository() GameRepository {
	return &gameRepository{
		storage: &sync.Map{},
	}
}

func (repo *gameRepository) SaveGame(game *model.Game) error {
	if game == nil {
		return errors.New("game cannot be nil")
	}

	// преобразуем domain к DTO
	gameDTO := mapper.ToDatasource(game)

	repo.storage.Store(game.UUID, gameDTO)
	return nil
}

func (repo *gameRepository) GetGame(id string) (*model.Game, error) {
	value, ok := repo.storage.Load(id)
	if !ok {
		return nil, errors.New("game not found")
	}

	gameDTO, ok := value.(*dto.GameDTO) // преобразуем обратно
	if !ok {
		return nil, errors.New("invalid data in storage")
	}

	return mapper.ToDomain(gameDTO), nil
}
