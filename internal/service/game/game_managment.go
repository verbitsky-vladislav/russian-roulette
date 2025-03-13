package game

import (
	"context"
	"errors"
	"fmt"
	"russian-roulette/internal/entities/custom_errors"
	gameEntities "russian-roulette/internal/entities/game"
	projectUtils "russian-roulette/internal/utils"
)

func (g *GameService) CreateGame(ctx context.Context, newGame *gameEntities.CreateGame) (*gameEntities.Game, error) {
	gameExists, err := g.gameRepository.GetAll(ctx, &gameEntities.GetGameFilters{
		CreatorUuid: &newGame.CreatorUuid,
		Status:      (*string)(&newGame.Status),
	})
	if err != nil {
		return nil, err
	}
	if len(gameExists) > 0 {
		return nil, errors.New(custom_errors.ErrGameAlreadyExists)
	}

	game, err := g.gameRepository.Create(ctx, newGame)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (g *GameService) CancelGame(ctx context.Context, gameUuid, creatorUuid string) error {
	game, err := g.gameRepository.GetAll(ctx, &gameEntities.GetGameFilters{
		Uuid:        projectUtils.ToPtr(gameUuid),
		CreatorUuid: projectUtils.ToPtr(creatorUuid),
	})
	if err != nil {
		return err
	}
	if len(game) == 0 {
		return errors.New(custom_errors.ErrGameNotFound)
	}

	_, err = g.gameRepository.Update(ctx, &gameEntities.UpdateGame{
		Uuid:   gameUuid,
		Status: projectUtils.ToPtr(gameEntities.Canceled),
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *GameService) StartGame(ctx context.Context, gameUuid string) error {
	_, err := g.gameRepository.Update(ctx, &gameEntities.UpdateGame{
		Uuid:   gameUuid,
		Status: projectUtils.ToPtr(gameEntities.Active),
	})
	if err != nil {
		return err
	}

	// Получаем список игроков
	players, err := g.gamePlayerRepository.GetAll(ctx, &gameEntities.GetGamePlayersFilters{
		GameUuid: projectUtils.ToPtr(gameUuid),
	})
	if err != nil {
		return err
	}

	if len(players) == 0 {
		return errors.New("игроки не найдены")
	}

	firstPlayer := players[0]

	_, err = g.gameRoundRepository.Create(ctx, &gameEntities.CreateGameRound{
		GameUuid: gameUuid,
		UserUuid: firstPlayer.UserUuid,
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *GameService) createTurnsQueue(gameUuid string, players []gameEntities.GamePlayer) error {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	for _, p := range players {
		err := g.cacheService.PushToQueue(redisKey, p.UserUuid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GameService) getNextTurn(gameUuid string) (string, error) {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	nextPlayerUuid, err := g.cacheService.PopFromQueue(redisKey)
	if err != nil {
		return "", err
	}

	if nextPlayerUuid == "" {
		return "", errors.New("очередь ходов пуста")
	}

	err = g.cacheService.PushToQueue(redisKey, nextPlayerUuid)
	if err != nil {
		return "", err
	}

	return nextPlayerUuid, nil
}

func (g *GameService) removeFromQueue(gameUuid, userUuid string) error {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	return g.cacheService.RemoveFromQueue(redisKey, userUuid)
}

func (g *GameService) isGameOver(gameUuid string) (bool, string, error) {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	players, err := g.cacheService.GetQueueValues(redisKey)
	if err != nil {
		return false, "", err
	}

	if len(players) == 1 {
		return true, players[0], nil
	}

	return false, "", nil
}

func (g *GameService) CreateRound(ctx context.Context, gameUuid, userUuid string) (*gameEntities.GameRound, error) {
	round, err := g.gameRoundRepository.Create(ctx, &gameEntities.CreateGameRound{
		GameUuid: gameUuid,
		UserUuid: userUuid,
	})
	if err != nil {
		return nil, err
	}

	return round, nil
}
