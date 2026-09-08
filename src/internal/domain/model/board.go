package model

import (
	"errors"
)

type Field struct {
	Cells [3][3]int // 0 - пусто, 1 - х, 2 - о
}

// .x.
// o.x
// o..

// {0, 1, 0},
// {2, 0, 1},
// {2, 0, 0}

func NewField() *Field {
	return &Field{
		Cells: [3][3]int{},
	}
}

func (field *Field) SetCell(row, col, value int) error {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return errors.New("indexes must be from 0 to 2")
	}

	if value < 0 || value > 2 {
		return errors.New("value must be 0 (empty), 1 (x) or 2 (o)")
	}

	field.Cells[row][col] = value
	return nil
}

func (field *Field) GetWinner() int {
	cells := field.Cells

	for i := 0; i < 3; i++ {
		if cells[i][0] != 0 && cells[i][0] == cells[i][1] && cells[i][1] == cells[i][2] {
			return cells[i][0]
		}
	}

	for j := 0; j < 3; j++ {
		if cells[0][j] != 0 && cells[0][j] == cells[1][j] && cells[1][j] == cells[2][j] {
			return cells[0][j]
		}
	}

	if cells[0][0] != 0 && cells[0][0] == cells[1][1] && cells[1][1] == cells[2][2] {
		return cells[0][0]
	}
	if cells[0][2] != 0 && cells[0][2] == cells[1][1] && cells[1][1] == cells[2][0] {
		return cells[0][2]
	}

	return 0
}

func (field *Field) IsDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if field.Cells[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
