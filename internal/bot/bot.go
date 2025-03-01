package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot/handler"
	"russian-roulette/internal/config"
)

type Bot struct {
	commandsHandler handler.CommandsHandler

	bot    *tgbotapi.BotAPI
	cfg    *config.Config
	logger *zap.Logger
}

func New(
	commandsHandler handler.CommandsHandler,

	bot *tgbotapi.BotAPI,
	cfg *config.Config, logger *zap.Logger,
) *Bot {
	return &Bot{
		commandsHandler: commandsHandler,

		bot:    bot,
		cfg:    cfg,
		logger: logger,
	}
}

func (b *Bot) Init() {

	u := tgbotapi.NewUpdate(0)
	updates := b.bot.GetUpdatesChan(u)

	h := handler.New(
		b.commandsHandler,
		b.bot, b.cfg, b.logger,
	)

	for update := range updates {
		go h.Handle(update)
	}

}
