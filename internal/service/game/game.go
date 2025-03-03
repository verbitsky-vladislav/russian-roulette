package game

import (
	"context"
	"go.uber.org/zap"
	gameEntities "russian-roulette/internal/entities/game"
	"russian-roulette/internal/repository"
)

type GameService struct {
	gameRepository       repository.GameRepository
	gameRoundRepository  repository.GameRoundRepository
	gamePlayerRepository repository.GamePlayerRepository
	logger               *zap.Logger
}

func NewGameService(
	gameRepository repository.GameRepository,
	gameRoundRepository repository.GameRoundRepository,
	gamePlayerRepository repository.GamePlayerRepository,
	logger *zap.Logger,
) *GameService {
	return &GameService{
		gameRepository:       gameRepository,
		gameRoundRepository:  gameRoundRepository,
		gamePlayerRepository: gamePlayerRepository,
		logger:               logger,
	}
}

func (g *GameService) GetGameByUuid(ctx context.Context, gameUuid string, rounds, players bool) (*gameEntities.Game, []*gameEntities.GameRound, []*gameEntities.GamePlayer, error) {
	var game *gameEntities.Game
	var gameRounds []*gameEntities.GameRound
	var gamePlayers []*gameEntities.GamePlayer

	game, err := g.gameRepository.GetByUUID(ctx, gameUuid)
	if err != nil {
		return nil, nil, nil, err
	}

	if rounds {
		gameRounds, err = g.gameRoundRepository.GetAll(ctx, &gameEntities.GetGameRoundsFilters{
			GameUuid: &gameUuid,
		})
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if players {
		gamePlayers, err = g.gamePlayerRepository.GetAll(ctx, &gameEntities.GetGamePlayersFilters{
			GameUuid: &gameUuid,
		})
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return game, gameRounds, gamePlayers, nil
}
