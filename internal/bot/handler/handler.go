package handler

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot/custom_errors"
	telegramUtils "russian-roulette/internal/bot/utils"
	"russian-roulette/internal/config"
	projectUtils "russian-roulette/internal/utils"
)

type Handler struct {
	// handlers
	commandsHandler  CommandsHandler
	callbacksHandler CallbacksHandler

	bot    *tgbotapi.BotAPI
	cfg    *config.Config
	logger *zap.Logger
}

type CommandsHandler interface {
	CommandsRouter(message *tgbotapi.Message) error
}

type CallbacksHandler interface {
	CallbacksRouter(message *tgbotapi.CallbackQuery) error
}

func New(
	commandsHandler CommandsHandler,
	callbacksHandler CallbacksHandler,

	bot *tgbotapi.BotAPI, cfg *config.Config, logger *zap.Logger,
) *Handler {
	return &Handler{
		commandsHandler:  commandsHandler,
		callbacksHandler: callbacksHandler,

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

	switch {
	case update.Message != nil: // Обработка сообщений

		// Обработка команд
		if update.Message.IsCommand() {
			h.handleCommand(update.Message)
			return
		}
		return
	case update.CallbackQuery != nil:
		h.handleCallback(update.CallbackQuery)
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
			"error in handle command handler",
			zap.Int64("user_id", message.Chat.ID),
			zap.String("command", message.Command()),
			zap.Error(err),
		)

		h.logger.Info("Error type", zap.String("type", fmt.Sprintf("%T", err)))

		var customError custom_errors.CustomError
		if errors.As(err, &customError) { // ✅ Правильная проверка
			err = telegramUtils.SendMessage(h.bot, &telegramUtils.Message{
				ChatId:           message.Chat.ID,
				Text:             err.Error(),
				IsRemoveKeyboard: projectUtils.ToPtr(true),
				MessageType:      projectUtils.ToPtr(telegramUtils.Deletable),
			}, h.logger)
			return
		}
		return
	}
}

func (h *Handler) handleCallback(message *tgbotapi.CallbackQuery) {
	err := h.callbacksHandler.CallbacksRouter(message)
	if err != nil {
		h.logger.Error(
			"error in handle callback query handler",
			zap.Int64("user_id", message.Message.Chat.ID),
			zap.String("callback data", message.Data),
			zap.Error(err),
		)

		var customError custom_errors.CustomError
		if errors.As(err, &customError) { // ✅ Правильная проверка
			err = telegramUtils.SendMessage(h.bot, &telegramUtils.Message{
				ChatId:           message.Message.Chat.ID,
				Text:             err.Error(),
				IsRemoveKeyboard: projectUtils.ToPtr(true),
				MessageType:      projectUtils.ToPtr(telegramUtils.Deletable),
			}, h.logger)
			return
		}
		return
	}
}
