package service

// бизнес логика, реализация интерфейса сервиса

import (
	"src/internal/datasource/repository"
	"src/internal/domain/model"
	"src/internal/domain/validator"
)

type MinimaxService struct {
	repo      repository.GameRepository
	validator *validator.GameValidator
	ai        *AI
}

type AI struct{}

func NewMinimaxService(repo repository.GameRepository) *MinimaxService {
	return &MinimaxService{
		repo:      repo,
		validator: validator.NewGameValidator(),
		ai:        &AI{},
	}
}

func (s *MinimaxService) CreateGame() (*model.Game, error) {
	game := model.NewGame()
	err := s.repo.SaveGame(game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

// комп ходит после вычисления ходов в минимаксе
func (service *MinimaxService) MakeMove(gameID string, row, column int) (*model.Game, error) {
	game, err := service.repo.GetGame(gameID)
	if err != nil {
		return nil, err
	}

	if err := service.validator.ValidateGameOver(game.Field); err != nil {
		return nil, err
	}

	if err := service.validator.ValidateMove(game.Field, row, column); err != nil {
		return nil, err
	}

	game.Field.Cells[row][column] = 1 // ход игрок

	// проверка завершения игры после хода игрока
	if game.CheckGameOver() {
		service.repo.SaveGame(game)
		return game, nil
	}

	// ход ИИ
	aiRow, aiColumn := service.ai.BestMove(game.Field)
	game.Field.Cells[aiRow][aiColumn] = 2

	service.repo.SaveGame(game)

	return game, nil
}
func (ai *AI) BestMove(field model.Field) (int, int) {
	bestScore := -1000
	bestRow, bestCol := -1, -1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if field.Cells[i][j] == 0 {
				newField := field
				newField.Cells[i][j] = 2

				score := ai.minimax(newField, 0, false)

				if score > bestScore {
					bestScore = score
					bestRow, bestCol = i, j
				}
			}
		}
	}

	return bestRow, bestCol
}

func (ai *AI) minimax(field model.Field, depth int, isMax bool) int {
	winner := field.GetWinner()

	if winner == 2 {
		return 10 - depth
	}
	if winner == 1 {
		return -10 + depth
	}
	if field.IsDraw() {
		return 0
	}

	if isMax {
		best := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if field.Cells[i][j] == 0 {
					newField := field
					newField.Cells[i][j] = 2

					score := ai.minimax(newField, depth+1, false)
					best = max(best, score)
				}
			}
		}
		return best
	}

	best := 1000
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if field.Cells[i][j] == 0 {
				newField := field
				newField.Cells[i][j] = 1

				score := ai.minimax(newField, depth+1, true)
				best = min(best, score)
			}
		}
	}
	return best
}

func (service *MinimaxService) ValidateFieldGame(gameID string, row, column int) error {
	game, err := service.repo.GetGame(gameID)
	if err != nil {
		return err
	}

	if err := service.validator.ValidateGameOver(game.Field); err != nil {
		return err
	}

	err = service.validator.ValidateMove(game.Field, row, column)
	return err
}

func (service *MinimaxService) CheckGameOver(field model.Field) bool {
	winner := field.GetWinner()
	return winner != 0 || field.IsDraw()
}
