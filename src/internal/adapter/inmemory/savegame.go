package inmemory

import (
	"tictac/internal/domain"
)

func (rep *InMemory) SaveGame(game *domain.Game) error {
	model := mapToModel(game)
	rep.storage.Store(model.Uuid, model)
	return nil
}
