package commands

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"russian-roulette/internal/bot/custom_errors"
	"russian-roulette/internal/bot/middleware"
	projectCustomErrors "russian-roulette/internal/entities/custom_errors"
	userEntities "russian-roulette/internal/entities/user"
)

// todo сделать
func (cmd *Commands) Pull(ctx context.Context, message *tgbotapi.Message) error {
	u, ok := ctx.Value(middleware.UserContextKey).(*userEntities.User)
	if !ok {
		return errors.New(projectCustomErrors.ErrUserNotFoundInContext)
	}

	_, err := cmd.userService.GetUserActiveGame(ctx, u.Uuid)
	if err != nil {
		if err.Error() == projectCustomErrors.ErrGameNotFound {
			return custom_errors.ErrUserWithoutActiveGame
		}
		return err
	}

	return nil
}

// todo сделать
func (cmd *Commands) Pass(ctx context.Context, message *tgbotapi.Message) error {
	u, ok := ctx.Value(middleware.UserContextKey).(*userEntities.User)
	if !ok {
		return errors.New(projectCustomErrors.ErrUserNotFoundInContext)
	}

	_, err := cmd.userService.GetUserActiveGame(ctx, u.Uuid)
	if err != nil {
		if err.Error() == projectCustomErrors.ErrGameNotFound {
			return custom_errors.ErrUserWithoutActiveGame
		}
		return err
	}

	return nil
}
