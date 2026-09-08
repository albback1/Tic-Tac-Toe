package inmemory

import (
	"fmt"
	"tictac/internal/domain"

	"github.com/google/uuid"
)

func (rep *InMemory) LoadGame(uuid uuid.UUID) (*domain.Game, error) {
	value, ok := rep.storage.Load(uuid)

	if !ok {
		return nil, fmt.Errorf("Game not found")
	}

	model, ok := value.(*GameRep)

	if !ok {
		return nil, fmt.Errorf("invalid type stored in map")
	}

	return mapToDomain(model), nil
}
