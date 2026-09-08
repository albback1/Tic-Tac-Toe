package dto

import (
	"tictac/internal/domain"
)

func MapToDomain(req *RequestGame) *domain.Game {
	game := domain.NewGame(len(req.Field), req.Uuid)

	for i := 0; i < len(game.Field); i++ {
		for j := 0; j < len(game.Field[i]); j++ {
			game.Field[i][j] = domain.Cell(req.Field[i][j])
		}
	}
	return game
}

func MapToDTO(game *domain.Game) *ResponseGame {
	resp := NewResponseGame(len(game.Field), game.UUID)
	for i := 0; i < len(game.Field); i++ {
		for j := 0; j < len(game.Field[i]); j++ {
			resp.Field[i][j] = int(game.Field[i][j])
		}
	}
	return resp
}
