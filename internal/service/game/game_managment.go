package game

import (
	"context"
	"errors"
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
	//players, err := g.gamePlayerRepository.GetAll(ctx, &gameEntities.GetGamePlayersFilters{
	//	GameUuid: &game.Uuid,
	//})
	//if err != nil {
	//	return err
	//}
	//
	//err = g.GenerateRounds(game, players)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (g *GameService) CreateRound(ctx context.Context, gameUuid, userUuid string, action gameEntities.GameAction, result gameEntities.GameActionResult) (*gameEntities.GameRound, error) {
	round, err := g.gameRoundRepository.Create(ctx, &gameEntities.CreateGameRound{
		GameUuid: gameUuid,
		UserUuid: userUuid,
		Action:   action,
		Result:   result,
	})
	if err != nil {
		return nil, err
	}

	return round, nil
}
