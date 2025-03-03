package game

import (
	"context"
	gameEntities "russian-roulette/internal/entities/game"
)

func (g *GameService) AddUserToGame(ctx context.Context, userUuid, gameUuid string) (*gameEntities.GamePlayer, error) {
	player, err := g.gamePlayerRepository.Create(ctx, &gameEntities.CreateGamePlayer{
		GameUuid: gameUuid,
		UserUuid: userUuid,
	})
	if err != nil {
		return nil, err
	}

	return player, nil
}
