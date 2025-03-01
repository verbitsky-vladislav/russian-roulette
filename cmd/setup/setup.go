package setup

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot"
	"russian-roulette/internal/bot/handler/commands"
	"russian-roulette/internal/config"
	"russian-roulette/internal/repository"
)

func Setup(cfg *config.Config, logger *zap.Logger) {
	// ctx

	// init db
	db := repository.New(&cfg.Database, logger)
	defer func() {
		err := db.Close()
		if err != nil {
			return
		}
	}()

	// repositories

	// init telegram bot
	botInstance, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		panic(err)
	}

	// handlers
	commandsHandler := commands.NewCommandsHandler(botInstance, logger)

	// Инициализация сервиса телеграм бота
	b := bot.New(
		commandsHandler,

		botInstance, cfg, logger,
	)

	b.Init() // Телеграм бот крутится в бесконечном цикле, ожидая обновлений
}
