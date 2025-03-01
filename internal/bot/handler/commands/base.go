package commands

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"russian-roulette/internal/bot/text"
	telegramUtils "russian-roulette/internal/bot/utils"
	userEntities "russian-roulette/internal/entities/user"
)

func (cmd *Commands) Start(ctx context.Context, message *tgbotapi.Message) error {
	u, err := cmd.userService.RegisterUser(&userEntities.CreateUser{
		ChatId: message.Chat.ID,
		TgName: message.Chat.UserName,
	})
	if err != nil {
		return err
	}

	err = telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
		ChatId:      u.ChatId,
		Text:        text.StartMessage(),
		MessageType: nil,
	}, cmd.logger)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *Commands) Help() error {
	return nil
}
