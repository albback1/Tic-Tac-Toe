package usecase

// Стурктура сервиса
type GameServiceItml struct {
	Repo GameRepository
}

// Констурктор сервиса
func NewGameServiceItml(repo GameRepository) *GameServiceItml {
	return &GameServiceItml{Repo: repo}
}
