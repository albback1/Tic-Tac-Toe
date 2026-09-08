package usecase

import (
	"tictac/internal/domain"
)

func (serv *GameServiceItml) ValidateField(newGame *domain.Game) (bool, error) {
	repoGame, err := serv.Repo.LoadGame(newGame.UUID)
	if err != nil {
		return false, err
	}

	var countMisMatch int

	for i := 0; i < len(repoGame.Field); i++ {
		for j := 0; j < len(repoGame.Field[i]); j++ {
			if repoGame.Field[i][j] != newGame.Field[i][j] {
				countMisMatch++
			}
		}
	}

	if countMisMatch == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
