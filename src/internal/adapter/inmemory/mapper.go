package inmemory

import (
	"tictac/internal/domain"
)

func mapToModel(game *domain.Game) *GameRep {
	rep := NewGameRep(len(game.Field), game.UUID)

	for i := 0; i < len(rep.Field); i++ {
		for j := 0; j < len(rep.Field[i]); j++ {
			rep.Field[i][j] = int(game.Field[i][j])
		}
	}

	return rep
}

func mapToDomain(rep *GameRep) *domain.Game {
	game := domain.NewGame(len(rep.Field), rep.Uuid)
	for i := 0; i < len(rep.Field); i++ {
		for j := 0; j < len(rep.Field[i]); j++ {
			game.Field[i][j] = domain.Cell(rep.Field[i][j])
		}
	}
	return game
}
