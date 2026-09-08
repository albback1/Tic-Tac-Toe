package usecase

import (
	"tictac/internal/domain"
)

func (repo *GameServiceItml) GameOverCheck(game *domain.Game) domain.GameResult {
	field := game.Field
	xSizeField := len(field[0])
	ySizeField := len(field)
	var result domain.GameResult = domain.Draw

	var xWon int = 0
	var oWon int = 0

	// проверяем горизонтальные значения
	for i := 0; i < ySizeField; i++ {
		for j := 0; j < xSizeField; j++ {
			if field[i][j] == domain.X {
				xWon++
			} else if field[i][j] == domain.O {
				oWon++
			} else {
				result = domain.Ongoing // вслучаи если есть пустые строки значит игра еще может продолжаться
			}
		}
		if xWon == xSizeField {
			return domain.XWin
		} else if oWon == xSizeField {
			return domain.OWin
		}
		xWon = 0
		oWon = 0
	}

	// проверяем вертикальные значения
	for i := 0; i < xSizeField; i++ {
		for j := 0; j < ySizeField; j++ {
			if field[j][i] == domain.X {
				xWon++
			} else if field[j][i] == domain.O {
				oWon++
			}
		}
		if xWon == xSizeField {
			return domain.XWin
		} else if oWon == xSizeField {
			return domain.OWin
		}
		xWon = 0
		oWon = 0
	}

	//проверяем диагональные значения слева
	for i := 0; i < ySizeField; i++ {
		if field[i][0+i] == domain.X {
			xWon++
		} else if field[i][0+i] == domain.O {
			oWon++
		}
	}

	if xWon == xSizeField {
		return domain.XWin
	} else if oWon == xSizeField {
		return domain.OWin
	}

	xWon = 0
	oWon = 0

	//проверяем диагональные значения справа
	for i := 0; i < ySizeField; i++ {
		if field[i][xSizeField-1-i] == domain.X {
			xWon++
		} else if field[i][xSizeField-1-i] == domain.O {
			oWon++
		}
	}

	if xWon == xSizeField {
		return domain.XWin
	} else if oWon == xSizeField {
		return domain.OWin
	}

	return result
}
