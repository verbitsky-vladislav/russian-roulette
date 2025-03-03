package user

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"russian-roulette/internal/entities/custom_errors"
	gameEntities "russian-roulette/internal/entities/game"
	userEntities "russian-roulette/internal/entities/user"
	"russian-roulette/internal/repository"
	"russian-roulette/internal/service"
)

type UserService struct {
	userRepo    repository.UserRepository
	gameService service.GameService
	logger      *zap.Logger
}

func NewUserService(userRepo repository.UserRepository, gameService service.GameService, logger *zap.Logger) *UserService {
	return &UserService{
		userRepo:    userRepo,
		gameService: gameService,
		logger:      logger,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, newUser *userEntities.CreateUser) (*userEntities.User, error) {
	u, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" { // user already exists
				u, err = s.userRepo.GetByChatID(ctx, newUser.ChatId)
				return u, nil
			}
		}
		return nil, err
	}

	return u, nil
}

func (s *UserService) GetUserByChatId(ctx context.Context, chatId int64) (*userEntities.User, error) {
	u, err := s.userRepo.GetByChatID(ctx, chatId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) JoinGame(ctx context.Context, userUuid, gameUuid string) error {
	s.logger.Info("JoinGame called", zap.String("userUuid", userUuid), zap.String("gameUuid", gameUuid))

	game, _, players, err := s.gameService.GetGameByUuid(ctx, gameUuid, true, true)
	if err != nil {
		s.logger.Error("Error fetching game by UUID", zap.Error(err))
		return err
	}
	if game == nil {
		s.logger.Warn("Game not found", zap.String("gameUuid", gameUuid))
		return errors.New(custom_errors.ErrGameNotFound)
	}

	s.logger.Info("Game found", zap.String("gameUuid", gameUuid), zap.Int("bulletCount", game.BulletCount), zap.Int("playersCount", len(players)), zap.String("status", string(game.Status)))

	if len(players) >= game.BulletCount+1 || game.Status == gameEntities.Active {
		s.logger.Warn("Game is already full or active", zap.String("gameUuid", gameUuid))
		return errors.New(custom_errors.ErrGameIsAlreadyFull)
	}

	for _, player := range players {
		if player.UserUuid == userUuid {
			s.logger.Warn("User already joined game", zap.String("userUuid", userUuid), zap.String("gameUuid", gameUuid))
			return errors.New(custom_errors.ErrUserAlreadyJoinToGame)
		}
	}

	s.logger.Info("Adding user to game", zap.String("userUuid", userUuid), zap.String("gameUuid", gameUuid))
	_, err = s.gameService.AddUserToGame(ctx, userUuid, gameUuid)
	if err != nil {
		s.logger.Error("Error adding user to game", zap.Error(err))
		return err
	}

	s.logger.Info("User successfully joined game", zap.String("userUuid", userUuid), zap.String("gameUuid", gameUuid))
	return nil
}
