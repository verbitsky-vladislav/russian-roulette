package middleware

import (
	"context"
	"database/sql"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"russian-roulette/internal/bot/custom_errors"
	userEntities "russian-roulette/internal/entities/user"
	"russian-roulette/internal/service"
)

func AuthMessageMiddleware(userService service.UserService) func(MessageHandlerFunc) MessageHandlerFunc {
	return func(next MessageHandlerFunc) MessageHandlerFunc {
		return func(ctx context.Context, message *tgbotapi.Message) error {
			// Проверяем, есть ли уже пользователь в контексте
			if user, ok := ctx.Value(UserContextKey).(*userEntities.User); ok && user != nil {
				return next(ctx, message)
			}

			// Получаем пользователя из базы
			user, err := userService.GetUserByChatId(ctx, message.From.ID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return custom_errors.ErrUserNotFound
				}
				return err
			}
			if user == nil {
				return custom_errors.ErrUserNotFound
			}

			// Добавляем пользователя в контекст
			ctx = context.WithValue(ctx, UserContextKey, user)

			// Передаем дальше по цепочке
			return next(ctx, message)
		}
	}
}

func AuthCallbackMiddleware(userService service.UserService) func(CallbackHandlerFunc) CallbackHandlerFunc {
	return func(next CallbackHandlerFunc) CallbackHandlerFunc {
		return func(ctx context.Context, message *tgbotapi.CallbackQuery) error {
			// Проверяем, есть ли уже пользователь в контексте
			if user, ok := ctx.Value(UserContextKey).(*userEntities.User); ok && user != nil {
				return next(ctx, message)
			}

			// Получаем пользователя из базы
			user, err := userService.GetUserByChatId(ctx, message.From.ID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return custom_errors.ErrUserNotFound
				}
				return err
			}
			if user == nil {
				return custom_errors.ErrUserNotFound
			}

			// Добавляем пользователя в контекст
			ctx = context.WithValue(ctx, UserContextKey, user)

			// Передаем дальше по цепочке
			return next(ctx, message)
		}
	}
}
