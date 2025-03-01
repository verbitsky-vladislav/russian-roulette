package commands

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cmd *Commands) Roulette(ctx context.Context, message *tgbotapi.Message) error {
	return nil
}

func (cmd *Commands) Join(ctx context.Context, message *tgbotapi.Message) error {
	return nil
}

func (cmd *Commands) Players() error {
	return nil
}
