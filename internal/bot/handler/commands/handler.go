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
	gameService service.GameService
	bot         *tgbotapi.BotAPI
	logger      *zap.Logger
}

func NewCommandsHandler(
	bot *tgbotapi.BotAPI,
	userService service.UserService,
	gameService service.GameService,
	logger *zap.Logger,
) *Commands {
	return &Commands{
		bot:         bot,
		userService: userService,
		gameService: gameService,
		logger:      logger,
	}
}

func (cmd *Commands) CommandsRouter(message *tgbotapi.Message) error {
	ctx := context.Background()

	rouletteHandler := middleware.ApplyMessageMiddlewares(cmd.Roulette, middleware.AuthMessageMiddleware(cmd.userService))
	pullHandler := middleware.ApplyMessageMiddlewares(cmd.Pull, middleware.AuthMessageMiddleware(cmd.userService))
	passHandler := middleware.ApplyMessageMiddlewares(cmd.Pass, middleware.AuthMessageMiddleware(cmd.userService))
	statsHandler := middleware.ApplyMessageMiddlewares(cmd.Stats, middleware.AuthMessageMiddleware(cmd.userService))

	// todo добавить сохранение всех сообщений об игре в массив и их удаление (кроме finish) после того как игра закончится / закенселится / будет длиться слишком долго (время экспирации 1 день)
	// todo все сообщения бота переделать на reply to user message (чтобы было понятнее с кем общается бот)
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
	case "pull":
		{
			cmd.logger.Debug("/pull : command was handled")
			// Игрок тянет курок (стреляет).
			// Бот случайным образом определяет, произошел ли выстрел.
			err := pullHandler(ctx, message)
			return err
		} // todo добавить валидацию ошибок (action not allowed, не твой ход и т.д.)
	case "pass":
		{
			cmd.logger.Debug("/pass : command was handled")
			// Игрок передает револьвер следующему участнику.
			// Бот должен проверить, сделал ли игрок хотя бы один выстрел перед передачей.
			err := passHandler(ctx, message)
			return err
		} // todo добавить валидацию ошибок (action not allowed, не твой ход и т.д.)
	case "players":
		{
			cmd.logger.Debug("/players : command was handled")
			// Показывает список игроков, участвующих в текущей игре.
			err := cmd.Players()
			return err
		}
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
