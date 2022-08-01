package api

// GameService contains the methods of the game service
type GameService interface{}

// GameRepository is what lets our service do db operations without knowing anything about the implementation
type GameRepository interface{}

type gameService struct {
    storage GameRepository
}

func NewGameService(gameRepo GameRepository) GameService {
    return &gameService{
        storage: gameRepo,
    }
}