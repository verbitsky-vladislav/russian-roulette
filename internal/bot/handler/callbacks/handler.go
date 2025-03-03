package callbacks

import (
	"context"
	"github.com/friendsofgo/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot/buttons/inline"
	botCustomErrors "russian-roulette/internal/bot/custom_errors"
	"russian-roulette/internal/bot/middleware"
	"russian-roulette/internal/bot/text"
	telegramUtils "russian-roulette/internal/bot/utils"
	projectCustomErrors "russian-roulette/internal/entities/custom_errors"
	userEntities "russian-roulette/internal/entities/user"
	"russian-roulette/internal/service"
	projectUtils "russian-roulette/internal/utils"
)

type Callbacks struct {
	bot         *tgbotapi.BotAPI
	userService service.UserService
	gameService service.GameService
	logger      *zap.Logger
}

func NewCallbacksHandler(bot *tgbotapi.BotAPI, userService service.UserService, gameService service.GameService, logger *zap.Logger) *Callbacks {
	return &Callbacks{
		bot:         bot,
		userService: userService,
		gameService: gameService,
		logger:      logger,
	}
}

func (cb *Callbacks) CallbacksRouter(message *tgbotapi.CallbackQuery) error {
	ctx := context.Background()

	joinHandler := middleware.ApplyCallbackMiddlewares(cb.Join, middleware.AuthCallbackMiddleware(cb.userService))
	cancelHandler := middleware.ApplyCallbackMiddlewares(cb.Cancel, middleware.AuthCallbackMiddleware(cb.userService))

	if inline.IsJoinGameCallback(message.Data) {
		err := joinHandler(ctx, message)
		if err != nil {
			return err
		}
	}

	if inline.IsCancelGameCallback(message.Data) {
		err := cancelHandler(ctx, message)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cb *Callbacks) Join(ctx context.Context, message *tgbotapi.CallbackQuery) error {

	u, ok := ctx.Value(middleware.UserContextKey).(*userEntities.User)
	if !ok {
		return errors.New(projectCustomErrors.ErrUserNotFoundInContext)
	}

	gameUuid, ok := inline.ExtractGameUuid(message.Data)
	if !ok {
		cb.logger.Warn("Game UUID extraction failed", zap.String("data", message.Data))
		return botCustomErrors.ErrGameIsNotActive
	}

	err := cb.userService.JoinGame(ctx, u.Uuid, gameUuid)
	if err != nil {
		errText := text.DefaultErrorMessage()

		if err.Error() == projectCustomErrors.ErrGameNotFound {
			errText = text.GameNotFoundMessage()
		}
		if err.Error() == projectCustomErrors.ErrGameIsAlreadyFull {
			errText = text.GameIsAlreadyFullMessage()
		}
		if err.Error() == projectCustomErrors.ErrUserAlreadyJoinToGame {
			errText = text.UserAlreadyJoinedMessage()
		}

		err = telegramUtils.SendMessage(cb.bot, &telegramUtils.Message{
			ChatId:      message.Message.Chat.ID,
			Text:        errText,
			MessageType: projectUtils.ToPtr(telegramUtils.Deletable),
		}, cb.logger)
		if err != nil {
			return err
		}
		return err
	}

	err = telegramUtils.SendMessage(cb.bot, &telegramUtils.Message{
		ChatId:      message.Message.Chat.ID,
		Text:        text.SuccessfulJoinGameMessage(),
		MessageType: projectUtils.ToPtr(telegramUtils.Deletable),
	}, cb.logger)
	if err != nil {
		return err
	}

	return nil
}

func (cb *Callbacks) Cancel(ctx context.Context, message *tgbotapi.CallbackQuery) error {
	u, ok := ctx.Value(middleware.UserContextKey).(*userEntities.User)
	if !ok {
		return errors.New(projectCustomErrors.ErrUserNotFoundInContext)
	}

	gameUuid, ok := inline.ExtractGameUuid(message.Data)
	if !ok {
		return botCustomErrors.ErrGameIsNotActive
	}

	err := cb.gameService.CancelGame(ctx, gameUuid, u.Uuid)
	if err != nil {
		return err
	}

	err = telegramUtils.SendMessage(cb.bot, &telegramUtils.Message{
		ChatId:      message.Message.Chat.ID,
		Text:        text.SuccessfulCancelGameMessage(),
		MessageType: projectUtils.ToPtr(telegramUtils.Deletable),
	}, cb.logger)
	telegramUtils.DeleteCallback(cb.bot, message, cb.logger)
	if err != nil {
		return err
	}

	return nil
}
