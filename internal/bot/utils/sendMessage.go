package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot/custom_errors"
	"time"
)

type MessageType string

const (
	Deletable    MessageType = "deletable"
	NonDeletable MessageType = "nondeletable"
)

type Message struct {
	ChatId           int64                              `json:"chat_id"`
	Text             string                             `json:"text"`
	ParseMode        *string                            `json:"parse_mode"`
	Keyboard         *[][]tgbotapi.KeyboardButton       `json:"keyboard"`
	InlineKeyboard   *[][]tgbotapi.InlineKeyboardButton `json:"inline_keyboard"`
	IsRemoveKeyboard *bool                              `json:"is_remove_keyboard"`
	MessageType      *MessageType                       `json:"message_type"`
}

// SendMessage отправляет сообщение через Telegram API
func SendMessage(b *tgbotapi.BotAPI, m *Message, logger *zap.Logger) error {
	msg := tgbotapi.NewMessage(m.ChatId, m.Text)

	// Устанавливаем режим парсинга сообщений, по дефолту markdown
	msg.ParseMode = "markdown"
	if m.ParseMode != nil {
		msg.ParseMode = *m.ParseMode
	}

	// Добавляем обычную клавиатуру, если передана
	if m.Keyboard != nil {
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
			Keyboard:        *m.Keyboard,
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
		}
	}

	// Добавляем inline-клавиатуру, если передана
	if m.InlineKeyboard != nil {
		msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
			InlineKeyboard: *m.InlineKeyboard,
		}
	}

	// Удаляем клавиатуру, если указано
	if m.IsRemoveKeyboard != nil && *m.IsRemoveKeyboard {
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	// Отправляем сообщение
	sentMsg, err := b.Send(msg)
	if err != nil {
		logger.Warn("Error sending system message to user", zap.Error(err))
		return custom_errors.ErrMessageSending
	}

	if m.MessageType != nil && *m.MessageType == Deletable {
		go func() {
			time.Sleep(7 * time.Second)
			deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
			_, err = b.Send(deleteMsg)
			if err != nil {
				logger.Warn("Error deleting old system message", zap.Error(err))
				return
			}
			return
		}()
	}

	return nil
}
