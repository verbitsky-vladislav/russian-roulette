package setup

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot"
	"russian-roulette/internal/bot/handler/commands"
	"russian-roulette/internal/config"
	"russian-roulette/internal/repository"
	userRepository "russian-roulette/internal/repository/user"
	userService "russian-roulette/internal/service/user"
)

func Setup(cfg *config.Config, logger *zap.Logger) {
	// init Ethereum clients
	//_, err := reader.NewEthereumReader(cfg.Blockchain.RPCURL, cfg.Blockchain.ContractAddress)
	//if err != nil {
	//	logger.Fatal("Failed to initialize Ethereum reader", zap.Error(err))
	//}
	//_, err = writer.NewEthereumWriter(cfg.Blockchain.RPCURL, cfg.Blockchain.ContractAddress, cfg.Blockchain.PrivateKey)
	//if err != nil {
	//	logger.Fatal("Failed to initialize Ethereum writer", zap.Error(err))
	//}

	// init db
	db := repository.New(&cfg.Database, logger)
	defer func() {
		err := db.Close()
		if err != nil {
			return
		}
	}()

	// repositories
	userRepo := userRepository.NewUserRepository(db.DB, logger)

	// service
	userUc := userService.NewUserService(userRepo, logger)

	// init telegram bot
	botInstance, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		panic(err)
	}

	// handlers
	commandsHandler := commands.NewCommandsHandler(botInstance, userUc, logger)

	// Инициализация сервиса телеграм бота
	b := bot.New(
		commandsHandler,

		botInstance, cfg, logger,
	)

	b.Init() // Телеграм бот крутится в бесконечном цикле, ожидая обновлений
}
