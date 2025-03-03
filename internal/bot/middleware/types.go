package middleware

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ContextKey string

const UserContextKey ContextKey = "user"

type MessageHandlerFunc func(ctx context.Context, message *tgbotapi.Message) error
type CallbackHandlerFunc func(ctx context.Context, message *tgbotapi.CallbackQuery) error

func ApplyMessageMiddlewares(handler MessageHandlerFunc, middlewares ...func(MessageHandlerFunc) MessageHandlerFunc) MessageHandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func ApplyCallbackMiddlewares(handler CallbackHandlerFunc, middlewares ...func(CallbackHandlerFunc) CallbackHandlerFunc) CallbackHandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
