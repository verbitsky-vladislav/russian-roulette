package game

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
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

// todo отменить можно только если игра еще не началась
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

func (g *GameService) StartGame(ctx context.Context, gameUuid string) (firstPlayer *gameEntities.GamePlayer, err error) {
	_, err = g.gameRepository.Update(ctx, &gameEntities.UpdateGame{
		Uuid:   gameUuid,
		Status: projectUtils.ToPtr(gameEntities.Active),
	})
	if err != nil {
		return nil, err
	}

	players, err := g.gamePlayerRepository.GetAll(ctx, &gameEntities.GetGamePlayersFilters{
		GameUuid: projectUtils.ToPtr(gameUuid),
	})
	if err != nil {
		return nil, err
	}

	if len(players) == 0 {
		return nil, errors.New("игроки не найдены")
	}

	// Перемешиваем порядок игроков случайным образом
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	// Сохраняем порядок ходов в Redis
	turnsKey := fmt.Sprintf("game:%s:turns", gameUuid)
	for _, player := range players {
		err := g.cacheService.PushToQueue(turnsKey, player.UserUuid)
		if err != nil {
			return nil, err
		}
	}

	// Устанавливаем первого игрока
	firstPlayer = players[0]
	_, err = g.gameRoundRepository.Create(ctx, &gameEntities.CreateGameRound{
		GameUuid: gameUuid,
		UserUuid: firstPlayer.UserUuid,
	})
	if err != nil {
		return nil, err
	}

	return firstPlayer, nil
}

// todo добавить распределение выигрыша, очистку redis кеша и прочее
func (g *GameService) FinishGame(ctx context.Context, gameUuid, winnerUuid string) error {
	_, err := g.gameRepository.Update(ctx, &gameEntities.UpdateGame{
		Uuid:   gameUuid,
		Status: projectUtils.ToPtr(gameEntities.Finished),
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *GameService) CreateRound(ctx context.Context, createRound *gameEntities.CreateGameRound) (*gameEntities.GameRound, error) {
	round, err := g.gameRoundRepository.Create(ctx, createRound)
	if err != nil {
		return nil, err
	}

	return round, nil
}

func (g *GameService) GetLastRound(ctx context.Context, gameUuid string) (*gameEntities.GameRound, error) {
	rounds, err := g.gameRoundRepository.GetAll(ctx, &gameEntities.GetGameRoundsFilters{
		GameUuid: projectUtils.ToPtr(gameUuid),
		Limit:    projectUtils.ToPtr(1),                 // Ограничиваем одним последним раундом
		OrderBy:  projectUtils.ToPtr("created_at DESC"), // Берем самый последний
	})
	if err != nil {
		return nil, err
	}

	if len(rounds) == 0 {
		return nil, nil
	}

	return rounds[0], nil
}
