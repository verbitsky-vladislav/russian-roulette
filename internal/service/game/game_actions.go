package game

import (
	"context"
	gameEntities "russian-roulette/internal/entities/game"
)

func (g *GameService) AddUserToGame(ctx context.Context, userUuid, gameUuid, name string) (*gameEntities.GamePlayer, error) {
	player, err := g.gamePlayerRepository.Create(ctx, &gameEntities.CreateGamePlayer{
		Name:     name,
		GameUuid: gameUuid,
		UserUuid: userUuid,
	})
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (g *GameService) Pull(ctx context.Context) {

}

func (g *GameService) Pass(ctx context.Context) {

}
