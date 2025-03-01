package setup

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot"
	"russian-roulette/internal/bot/handler/commands"
	"russian-roulette/internal/config"
	"russian-roulette/internal/repository"
	"russian-roulette/internal/repository/reader"
	"russian-roulette/internal/repository/writer"
)

func Setup(cfg *config.Config, logger *zap.Logger) {
	// ctx
	//ctx := context.Background()

	// init Ethereum clients
	ethReader, err := reader.NewEthereumReader(cfg.Blockchain.RPCURL, cfg.Blockchain.ContractAddress)
	if err != nil {
		logger.Fatal("Failed to initialize Ethereum reader", zap.Error(err))
	}

	owner, err := ethReader.Owner(context.Background())
	if err != nil {
		logger.Fatal("Failed to fetch contract owner", zap.Error(err))
	}
	logger.Info("Contract Owner", zap.String("owner", owner.Hex()))

	minBet, err := ethReader.MinimumBet(context.Background())
	if err != nil {
		logger.Fatal("Failed to fetch minimum bet", zap.Error(err))
	}
	logger.Info("Minimum Bet", zap.String("minBet", minBet.String()))

	_, _ = writer.NewEthereumWriter(cfg.Blockchain.RPCURL, cfg.Blockchain.ContractAddress, cfg.Blockchain.PrivateKey)
	if err != nil {
		logger.Fatal("Failed to initialize Ethereum writer", zap.Error(err))
	}

	// init db
	db := repository.New(&cfg.Database, logger)
	defer func() {
		err := db.Close()
		if err != nil {
			return
		}
	}()

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
