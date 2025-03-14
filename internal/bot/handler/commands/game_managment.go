package commands

import (
	"context"
	"github.com/friendsofgo/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"russian-roulette/internal/bot/buttons/inline"
	"russian-roulette/internal/bot/middleware"
	"russian-roulette/internal/bot/text"
	telegramUtils "russian-roulette/internal/bot/utils"
	"russian-roulette/internal/entities/custom_errors"
	gameEntities "russian-roulette/internal/entities/game"
	"russian-roulette/internal/entities/types"
	userEntities "russian-roulette/internal/entities/user"
	projectUtils "russian-roulette/internal/utils"
	"strconv"
	"strings"
)

// todo fix messages
func (cmd *Commands) Roulette(ctx context.Context, message *tgbotapi.Message) error {
	u, ok := ctx.Value(middleware.UserContextKey).(*userEntities.User)
	if !ok {
		return errors.New(custom_errors.ErrUserNotFoundInContext)
	}

	args := strings.Fields(message.Text) // Используем Fields, чтобы убрать лишние пробелы

	if len(args) != 3 {
		err := telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
			ChatId:      message.Chat.ID,
			Text:        text.WrongRouletteCommandMessage(),
			MessageType: projectUtils.ToPtr(telegramUtils.Deletable),
		}, cmd.logger)
		telegramUtils.DeleteMessage(cmd.bot, message, cmd.logger)
		if err != nil {
			return err
		}

		return err
	}

	players, err := strconv.Atoi(args[1])
	bet := types.NewDecimalFromString(args[2])
	if err != nil || players < 2 || players > 6 || bet.Cmp(types.NewDecimalFromInt(0).Big) <= 0 {
		err := telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
			ChatId:      message.Chat.ID,
			Text:        text.WrongRouletteParamsMessage(),
			MessageType: projectUtils.ToPtr(telegramUtils.Deletable),
		}, cmd.logger)
		telegramUtils.DeleteMessage(cmd.bot, message, cmd.logger)
		return err
	}

	bullets := players - 1

	g, err := cmd.gameService.CreateGame(ctx, &gameEntities.CreateGame{
		CreatorUuid: u.Uuid,
		Status:      gameEntities.Waiting,
		BetAmount:   *bet,
		BulletCount: bullets,
		RoundsCount: 6, // todo maybe here need not constant
	})
	if err != nil {
		if err.Error() == custom_errors.ErrGameAlreadyExists {
			err := telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
				ChatId:      message.Chat.ID,
				Text:        text.GameAlreadyExistsMessage(),
				MessageType: projectUtils.ToPtr(telegramUtils.Deletable),
			}, cmd.logger)
			telegramUtils.DeleteMessage(cmd.bot, message, cmd.logger)
			return err
		}
		return err
	}

	err = telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
		ChatId:         message.Chat.ID,
		Text:           text.NewRouletteGameMessage(players, bullets, bet),
		InlineKeyboard: projectUtils.ToPtr(inline.CreateJoinGameButton(g.Uuid)),
	}, cmd.logger)

	return err
}

func (cmd *Commands) Players() error {
	return nil
}
