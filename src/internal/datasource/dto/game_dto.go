package dto

// шаблон для обмена данными между клиентом и API
// DTO = Data Transfer Object

import "time"

type GameDTO struct {
	ID        string
	Cells     [3][3]int
	CreatedAt time.Time
	UpdatedAt time.Time
}
