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
	ctx      context.Context
}

func NewUserService(userRepo repository.UserRepository, logger *zap.Logger, ctx context.Context) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
		ctx:      ctx,
	}
}

func (s *UserService) RegisterUser(newUser *userEntities.CreateUser) (*userEntities.User, error) {
	u, err := s.userRepo.Create(s.ctx, newUser)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" { // user already exists
				u, err = s.userRepo.GetByChatID(s.ctx, newUser.ChatId)
				return u, nil
			}
		}
		return nil, err
	}

	return u, nil
}

func (s *UserService) GetUserByChatId(chatId int64) (*userEntities.User, error) {
	u, err := s.userRepo.GetByChatID(s.ctx, chatId)
	if err != nil {
		return nil, err
	}

	return u, nil
}
