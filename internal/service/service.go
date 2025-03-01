package service

import (
	userEntities "russian-roulette/internal/entities/user"
)

type (
	UserService interface {
		RegisterUser(newUser *userEntities.CreateUser) (*userEntities.User, error)
		GetUserByChatId(chatId int64) (*userEntities.User, error)
		GetAllUsers(filters *userEntities.GetUserFilters) ([]*userEntities.User, error)
		UpdateUser(updUser *userEntities.UpdateUser) (*userEntities.User, error)
	}
)
