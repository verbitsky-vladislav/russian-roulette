package commands

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot/custom_errors"
	"russian-roulette/internal/bot/middleware"
	telegramUtils "russian-roulette/internal/bot/utils"
	"russian-roulette/internal/service"
	projectUtils "russian-roulette/internal/utils"
)

type Commands struct {
	userService service.UserService
	bot         *tgbotapi.BotAPI
	logger      *zap.Logger
}

func NewCommandsHandler(bot *tgbotapi.BotAPI, userService service.UserService, logger *zap.Logger) *Commands {
	return &Commands{
		bot:         bot,
		userService: userService,
		logger:      logger,
	}
}

func (cmd *Commands) CommandsRouter(message *tgbotapi.Message) error {
	ctx := context.Background()

	rouletteHandler := middleware.ApplyMiddlewares(cmd.Roulette, middleware.AuthMiddleware(cmd.userService))
	joinHandler := middleware.ApplyMiddlewares(cmd.Join, middleware.AuthMiddleware(cmd.userService))
	betHandler := middleware.ApplyMiddlewares(cmd.Bet, middleware.AuthMiddleware(cmd.userService))
	pullHandler := middleware.ApplyMiddlewares(cmd.Pull, middleware.AuthMiddleware(cmd.userService))
	passHandler := middleware.ApplyMiddlewares(cmd.Pass, middleware.AuthMiddleware(cmd.userService))
	statsHandler := middleware.ApplyMiddlewares(cmd.Stats, middleware.AuthMiddleware(cmd.userService))

	switch message.Command() {
	case "start":
		{
			cmd.logger.Debug("/start : command was handled")
			// Команда для запуска бота.
			// Должна отправлять пользователю приветственное сообщение и правила игры.
			err := cmd.chatOnly(cmd.Start, ctx, message)
			return err
		}
	case "help":
		{
			cmd.logger.Debug("/help : command was handled")
			// Показывает список всех доступных команд и их описание.
			err := cmd.Help()
			return err
		}
	case "roulette":
		{
			cmd.logger.Debug("/roulette : command was handled")
			// Создает новую игру русской рулетки.
			// Формат: /roulette [кол-во игроков] [размер магазина] [ставка]
			// Бот должен проверять корректность параметров и создавать новую игровую сессию.
			err := rouletteHandler(ctx, message)
			return err
		}
	case "join":
		{
			cmd.logger.Debug("/join : command was handled")
			// Позволяет пользователю присоединиться к уже созданной игре.
			// Бот должен проверять, есть ли активная игра и свободные места.
			err := joinHandler(ctx, message)
			return err
		}
	case "bet":
		{
			cmd.logger.Debug("/bet : command was handled")
			// Игрок делает ставку перед началом игры (если ставки нужны отдельно).
			// Проверяет баланс игрока и вносит сумму в банк игры.
			err := betHandler(ctx, message)
			return err
		}
	case "pull":
		{
			cmd.logger.Debug("/pull : command was handled")
			// Игрок тянет курок (стреляет).
			// Бот случайным образом определяет, произошел ли выстрел.
			err := pullHandler(ctx, message)
			return err
		}
	case "pass":
		{
			cmd.logger.Debug("/pass : command was handled")
			// Игрок передает револьвер следующему участнику.
			// Бот должен проверить, сделал ли игрок хотя бы один выстрел перед передачей.
			err := passHandler(ctx, message)
			return err
		}
	case "players":
		{
			cmd.logger.Debug("/players : command was handled")
			// Показывает список игроков, участвующих в текущей игре.
			err := cmd.Players()
			return err
		}
	//case "balance":
	//	{
	//		cmd.logger.Debug("/balance : command was handled")
	//		// Отображает баланс игрока в игре (например, количество токенов или денег).
	//		err := cmd.Balance()
	//		return err
	//	}
	case "stats":
		{
			cmd.logger.Debug("/stats : command was handled")
			// Показывает статистику игрока: количество побед, поражений и общий выигрыш.
			err := statsHandler(ctx, message)
			return err
		}
	case "top":
		{
			cmd.logger.Debug("/top : command was handled")
			// Выводит таблицу лидеров с лучшими игроками по выживанию и выигрышам.
			err := cmd.Top()
			return err
		}
	//case "wallet":
	//	{
	//		cmd.logger.Debug("/wallet : command was handled")
	//		// Проверяет криптокошелек игрока, если игра использует токены.
	//		err := cmd.Wallet()
	//		return err
	//	}
	//case "deposit":
	//	{
	//		cmd.logger.Debug("/deposit : command was handled")
	//		// Позволяет пополнить баланс игрока (например, с криптокошелька).
	//		err := cmd.Deposit()
	//		return err
	//	}
	//case "withdraw":
	//	{
	//		cmd.logger.Debug("/withdraw : command was handled")
	//		// Позволяет вывести выигранные средства с баланса игрока.
	//		err := cmd.Withdraw()
	//		return err
	//	}
	default:
		return custom_errors.ErrNoCommandFound
	}
}

func (cmd *Commands) chatOnly(next func(ctx context.Context, message *tgbotapi.Message) error, ctx context.Context, message *tgbotapi.Message) error {
	if message.Chat.Type == "group" || message.Chat.Type == "supergroup" {

		err := telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
			ChatId:      message.Chat.ID,
			Text:        custom_errors.ErrChatOnlyCommand.Error(),
			MessageType: projectUtils.ToPtr(telegramUtils.Deletable),
			ParseMode:   projectUtils.ToPtr(""),
		}, cmd.logger)
		if err != nil {
			cmd.logger.Warn("chatOnly: failed to send error message", zap.Error(err))
			return err
		}

		telegramUtils.DeleteMessage(cmd.bot, message, cmd.logger)

		return nil
	}

	return next(ctx, message)
}

func (cmd *Commands) groupOnly(next func(ctx context.Context, message *tgbotapi.Message) error, ctx context.Context, message *tgbotapi.Message) error {
	if !(message.Chat.Type == "group") || !(message.Chat.Type == "supergroup") {
		err := telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
			ChatId:      message.Chat.ID,
			Text:        custom_errors.ErrGroupOnlyCommand.Error(),
			MessageType: projectUtils.ToPtr(telegramUtils.Deletable),
			ParseMode:   projectUtils.ToPtr(""),
		}, cmd.logger)
		if err != nil {
			cmd.logger.Warn("groupOnly: failed to send error message", zap.Error(err))
			return err
		}

		telegramUtils.DeleteMessage(cmd.bot, message, cmd.logger)

		return nil
	}

	return next(ctx, message)
}
