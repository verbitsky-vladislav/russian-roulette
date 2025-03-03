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
		//GetAllUsers(filters *userEntities.GetUserFilters) ([]*userEntities.User, error)
		//UpdateUser(updUser *userEntities.UpdateUser) (*userEntities.User, error)
		JoinGame(ctx context.Context, userUuid, gameUuid string) error
	}
	GameService interface {
		GetGameByUuid(ctx context.Context, gameUuid string, rounds, players bool) (*gameEntities.Game, []*gameEntities.GameRound, []*gameEntities.GamePlayer, error)

		CreateGame(ctx context.Context, newGame *gameEntities.CreateGame) (*gameEntities.Game, error)
		CancelGame(ctx context.Context, gameUuid, creatorUuid string) error

		AddUserToGame(ctx context.Context, userUuid, gameUuid string) (*gameEntities.GamePlayer, error)
	}
)
