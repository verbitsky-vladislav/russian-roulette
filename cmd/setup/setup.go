package setup

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot"
	"russian-roulette/internal/bot/handler/callbacks"
	"russian-roulette/internal/bot/handler/commands"
	"russian-roulette/internal/config"
	"russian-roulette/internal/repository"
	gameRepository "russian-roulette/internal/repository/game"
	userRepository "russian-roulette/internal/repository/user"
	"russian-roulette/internal/service/cache"
	gameService "russian-roulette/internal/service/game"
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
	gameRepo := gameRepository.NewGameRepository(db.DB, logger)
	gameRoundsRepo := gameRepository.NewGameRoundRepository(db.DB, logger)
	gamePlayersRepo := gameRepository.NewGamePlayersRepository(db.DB, logger)

	// service
	cacheService := cache.NewRedisCache(&cfg.Redis, logger)
	gameUc := gameService.NewGameService(gameRepo, gameRoundsRepo, gamePlayersRepo, cacheService, logger)
	userUc := userService.NewUserService(userRepo, gameUc, logger)

	// init telegram bot
	botInstance, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		panic(err)
	}

	// handlers
	commandsHandler := commands.NewCommandsHandler(botInstance, userUc, gameUc, logger)
	callbacksHandler := callbacks.NewCallbacksHandler(botInstance, userUc, gameUc, logger)

	// Инициализация сервиса телеграм бота
	b := bot.New(
		commandsHandler,
		callbacksHandler,

		botInstance, cfg, logger,
	)

	b.Init() // Телеграм бот крутится в бесконечном цикле, ожидая обновлений
}
