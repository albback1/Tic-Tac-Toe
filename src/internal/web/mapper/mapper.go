package mapper

// мапперы domain<->web

import (
	"src/internal/domain/model"
	"src/internal/web/dto"
)

// domain->web
func ToResponse(game *model.Game) *dto.GameResponse {
	return &dto.GameResponse{
		ID:       game.UUID,
		Field:    game.Field.Cells,
		Winner:   game.GetWinner(),
		GameOver: game.CheckGameOver(),
	}
}

// domain<-web
func ToField(cells [3][3]int) model.Field {
	field := model.NewField()
	field.Cells = cells
	return *field
}
