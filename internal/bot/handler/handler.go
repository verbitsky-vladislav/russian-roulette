package handler

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot/custom_errors"
	telegramUtils "russian-roulette/internal/bot/utils"
	"russian-roulette/internal/config"
	projectUtils "russian-roulette/internal/utils"
)

type Handler struct {
	// handlers
	commandsHandler CommandsHandler

	bot    *tgbotapi.BotAPI
	cfg    *config.Config
	logger *zap.Logger
}

type CommandsHandler interface {
	CommandsRouter(message *tgbotapi.Message) error
}

func New(
	commandsHandler CommandsHandler,

	bot *tgbotapi.BotAPI, cfg *config.Config, logger *zap.Logger,
) *Handler {
	return &Handler{
		commandsHandler: commandsHandler,

		bot:    bot,
		cfg:    cfg,
		logger: logger,
	}
}

// Handle функция отвечает за получение и роутинг сообщений. Так же функция обрабатывает все ошибки, поступающие из обработчиков
func (h *Handler) Handle(update tgbotapi.Update) {
	defer func() {
		if r := recover(); r != nil {
			h.logger.Error(
				"Recovered from panic in Handle",
				zap.Any("panic", r),
				zap.String("stacktrace", projectUtils.GetStackTrace()), // добавляем stack trace
			)
		}
	}()

	h.logger.Info("Received update", zap.Any("update", update)) // Логируем ВСЁ

	switch {
	case update.Message != nil: // Обработка сообщений

		// Обработка команд
		if update.Message.IsCommand() {
			h.handleCommand(update.Message)
			return
		}
		return
	case update.CallbackQuery != nil:
		return
	default:
		h.logger.Warn("Unknown update type", zap.Any("update", update))
		return
	}
}

// todo сделать обработку ошибок
func (h *Handler) handleCommand(message *tgbotapi.Message) {
	err := h.commandsHandler.CommandsRouter(message)
	if err != nil {
		h.logger.Error(
			"error in handle command handler query",
			zap.Int64("user_id", message.Chat.ID),
			zap.String("command", message.Command()),
			zap.Error(err),
		)

		var customError custom_errors.CustomError
		if errors.As(err, &customError) {
			err = telegramUtils.SendMessage(h.bot, &telegramUtils.Message{
				ChatId:           message.Chat.ID,
				Text:             err.Error(),
				IsRemoveKeyboard: projectUtils.ToPtr(true),
				MessageType:      projectUtils.ToPtr(telegramUtils.Deletable),
			}, h.logger)
			return
		}
		h.logger.Warn("error in command handler", zap.Error(err))
		return
	}
}
