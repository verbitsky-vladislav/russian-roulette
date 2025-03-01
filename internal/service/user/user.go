package user

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/lib/pq"
	"go.uber.org/zap"
	userEntities "russian-roulette/internal/entities/user"
	"russian-roulette/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
	logger   *zap.Logger
}

func NewUserService(userRepo repository.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
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
