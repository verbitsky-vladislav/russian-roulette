package service

import (
	"context"
	userEntities "russian-roulette/internal/entities/user"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, newUser *userEntities.CreateUser) (*userEntities.User, error)
		GetUserByChatId(ctx context.Context, chatId int64) (*userEntities.User, error)
		//GetAllUsers(filters *userEntities.GetUserFilters) ([]*userEntities.User, error)
		//UpdateUser(updUser *userEntities.UpdateUser) (*userEntities.User, error)
	}
)
