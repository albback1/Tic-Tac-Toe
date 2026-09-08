package mapper

// мапперы domain<->datasource

import (
	datasource "src/internal/datasource/dto"
	domain "src/internal/domain/model"
)

// domain->datasource
func ToDatasource(game *domain.Game) *datasource.GameDTO {
	return &datasource.GameDTO{
		ID:        game.UUID,
		Cells:     game.Field.Cells,
		CreatedAt: game.CreatedAt,
		UpdatedAt: game.UpdatedAt,
	}
}

// domain<-datasource
func ToDomain(dto *datasource.GameDTO) *domain.Game {
	field := domain.NewField()
	field.Cells = dto.Cells

	return &domain.Game{
		UUID:      dto.ID,
		Field:     *field,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
