package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"time"
)

func DeleteMessage(b *tgbotapi.BotAPI, m *tgbotapi.Message, logger *zap.Logger) {
	go func() {
		time.Sleep(7 * time.Second)
		deleteMsg := tgbotapi.NewDeleteMessage(m.Chat.ID, m.MessageID)

		_, err := b.Request(deleteMsg)
		if err != nil {
			return
		}

	}()
}

func DeleteCallback(b *tgbotapi.BotAPI, m *tgbotapi.CallbackQuery, logger *zap.Logger) {
	go func() {
		time.Sleep(7 * time.Second)
		deleteMsg := tgbotapi.NewDeleteMessage(m.Message.Chat.ID, m.Message.MessageID)

		_, err := b.Request(deleteMsg)
		if err != nil {
			return
		}

	}()
}
