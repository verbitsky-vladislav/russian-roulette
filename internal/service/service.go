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
		JoinGame(ctx context.Context, userUuid, gameUuid, name string) (isStart bool, firstPlayer *gameEntities.GamePlayer, players []*gameEntities.GamePlayer, err error)
		CheckUserActiveGame(ctx context.Context, userUuid string) (bool, error)
		GetUserActiveGame(ctx context.Context, userUuid string) (*gameEntities.Game, error)
	}
	GameService interface {
		GetGameByUuid(ctx context.Context, gameUuid string, rounds, players bool) (*gameEntities.Game, []*gameEntities.GameRound, []*gameEntities.GamePlayer, error)
		GetAllGames(ctx context.Context, filters *gameEntities.GetGameFilters) ([]*gameEntities.Game, error)

		CreateGame(ctx context.Context, newGame *gameEntities.CreateGame) (*gameEntities.Game, error)
		CancelGame(ctx context.Context, gameUuid, creatorUuid string) error
		StartGame(ctx context.Context, gameUuid string) (firstPlayer *gameEntities.GamePlayer, err error)
		CreateRound(ctx context.Context, createRound *gameEntities.CreateGameRound) (*gameEntities.GameRound, error)
		GetLastRound(ctx context.Context, gameUuid string) (*gameEntities.GameRound, error)

		PullTrigger(ctx context.Context, game *gameEntities.Game, playerGuid string) (isDead, isOver bool, currentPlayer, nextPlayer *gameEntities.GamePlayer, updatedGame *gameEntities.Game, err error)
		PassTrigger(ctx context.Context, gameUuid, userUuid string) (currentPlayer, nextPlayer *gameEntities.GamePlayer, bullets, rounds int, err error)

		AddUserToGame(ctx context.Context, userUuid, gameUuid, name string) (*gameEntities.GamePlayer, error)
	}
)
