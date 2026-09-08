package domain

type GameService interface {
	ValidateField(newGame *Game) (bool, error)          // возвращает булевое значение если true если игра не менялась false если были внесены изменения
	GameOverCheck(game *Game) GameResult                // проверяет состояние игры и возвращает его
	NextMoveMinimax(game *Game, player Cell) (x, y int) // возращает координаты нового хода компьютера
}
