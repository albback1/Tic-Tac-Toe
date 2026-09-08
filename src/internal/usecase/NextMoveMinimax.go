package usecase

import (
	"fmt"
	"tictac/internal/domain"
)

type move struct {
	Score int
	X     int
	Y     int
}

func (service *GameServiceItml) NextMoveMinimax(game *domain.Game, player domain.Cell) (x, y int) {
	var coord []move = availableMoves(game)  // возвращет список возможных ходов
	nxtPlayer, _ := nextPlayer(player)       //определяет следующего ходящего
	scoreCheck(game, coord, service, player) // проверяет результат хода и ставит соответсвующее значение
	var maxScore int = -100
	var resIndex int

	for i, x := range coord {
		if x.Score > maxScore {
			maxScore = x.Score
			resIndex = i
		}
	}

	for i := 0; i < len(coord); i++ {
		if coord[i].Score == -1 {
			bufferGame := cloneGame(game)
			bufferGame.Field[coord[i].Y][coord[i].X] = player
			coord[i].Score = miniMaxRecurs(bufferGame, nxtPlayer, service)
		}
	}

	for i, x := range coord {
		if x.Score > maxScore {
			maxScore = x.Score
			resIndex = i
		}
	}

	return coord[resIndex].X, coord[resIndex].Y
}

func miniMaxRecurs(game *domain.Game, player domain.Cell, service *GameServiceItml) int {
	var coord []move = availableMoves(game)
	nxtPlayer, _ := nextPlayer(player)
	var result int

	scoreCheck(game, coord, service, player) // проверяет результат хода и ставит соответсвующее значение
	result = choiceMove(coord, player)

	if result == -10 && player == domain.O || result == 10 && player == domain.X {
		return result
	}

	for i := 0; i < len(coord); i++ {
		if coord[i].Score == -1 {
			bufferGame := cloneGame(game)
			bufferGame.Field[coord[i].Y][coord[i].X] = player
			coord[i].Score = miniMaxRecurs(bufferGame, nxtPlayer, service)
		}
	}
	result = choiceMove(coord, player)

	return result
}

func choiceMove(coord []move, player domain.Cell) int {
	var result int = -100
	if player == domain.O {
		result = 100
	}

	if player == domain.X {
		for _, x := range coord {
			if x.Score > result {
				result = x.Score
			}
		}
	} else if player == domain.O {
		for _, x := range coord {
			if x.Score < result {
				result = x.Score
			}
		}
	}
	return result
}

func scoreCheck(game *domain.Game, coord []move, service *GameServiceItml, player domain.Cell) {
	for i := 0; i < len(coord); i++ {
		bufferGame := cloneGame(game)

		bufferGame.Field[coord[i].Y][coord[i].X] = player
		checkRes := service.GameOverCheck(bufferGame)
		switch checkRes {
		case domain.Ongoing:
			coord[i].Score = -1
		case domain.Draw:
			coord[i].Score = +0
		case domain.OWin:
			coord[i].Score = -10
		case domain.XWin:
			coord[i].Score = +10
		}
	}
}

func cloneGame(game *domain.Game) *domain.Game {
	newGame := domain.NewGame(len(game.Field), game.UUID)

	for i := 0; i < len(game.Field); i++ {
		for j := 0; j < len(game.Field[i]); j++ {
			newGame.Field[i][j] = game.Field[i][j]
		}
	}

	return newGame
}

func nextPlayer(player domain.Cell) (domain.Cell, error) {
	switch player {
	case domain.X:
		return domain.O, nil
	case domain.O:
		return domain.X, nil
	default:
		return player, fmt.Errorf("player is empty")
	}
}

func availableMoves(game *domain.Game) []move {
	var result []move
	for i := 0; i < len(game.Field); i++ {
		for j := 0; j < len(game.Field[i]); j++ {
			if game.Field[i][j] == domain.Empty {
				var newCoord move = move{-1, j, i}
				result = append(result, newCoord)
			}
		}
	}
	return result
}
