package commands

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cmd *Commands) Balance() error {
	return nil
}

func (cmd *Commands) Stats(ctx context.Context, message *tgbotapi.Message) error {
	return nil
}

func (cmd *Commands) Top() error {
	return nil
}
