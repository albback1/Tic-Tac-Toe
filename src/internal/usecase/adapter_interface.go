package usecase

import (
	"tictac/internal/domain"

	"github.com/google/uuid"
)

type GameRepository interface {
	SaveGame(game *domain.Game) error
	LoadGame(uuid uuid.UUID) (*domain.Game, error)
}
