package domain

type NotFoundGame struct{}

func (err *NotFoundGame) Error() string {
	return "Game not found"
}
