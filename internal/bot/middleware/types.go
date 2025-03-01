package middleware

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ContextKey string

const UserContextKey ContextKey = "user"

type HandlerFunc func(ctx context.Context, message *tgbotapi.Message) error

func ApplyMiddlewares(handler HandlerFunc, middlewares ...func(HandlerFunc) HandlerFunc) HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
