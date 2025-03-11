package service

import (
	"context"
	gameEntities "russian-roulette/internal/entities/game"
	userEntities "russian-roulette/internal/entities/user"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, newUser *userEntities.CreateUser) (*userEntities.User, error)
		GetUserByChatId(ctx context.Context, chatId int64) (*userEntities.User, error)
		JoinGame(ctx context.Context, userUuid, gameUuid, name string) (bool, []*gameEntities.GamePlayer, error)
	}
	GameService interface {
		GetGameByUuid(ctx context.Context, gameUuid string, rounds, players bool) (*gameEntities.Game, []*gameEntities.GameRound, []*gameEntities.GamePlayer, error)

		CreateGame(ctx context.Context, newGame *gameEntities.CreateGame) (*gameEntities.Game, error)
		CancelGame(ctx context.Context, gameUuid, creatorUuid string) error
		StartGame(ctx context.Context, gameUuid string) error
		CreateRound(ctx context.Context, gameUuid, userUuid string, action gameEntities.GameAction, result gameEntities.GameActionResult) (*gameEntities.GameRound, error)

		AddUserToGame(ctx context.Context, userUuid, gameUuid, name string) (*gameEntities.GamePlayer, error)
	}
)
