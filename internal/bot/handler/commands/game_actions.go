package commands

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"russian-roulette/internal/bot/custom_errors"
	"russian-roulette/internal/bot/middleware"
	"russian-roulette/internal/bot/text"
	telegramUtils "russian-roulette/internal/bot/utils"
	projectCustomErrors "russian-roulette/internal/entities/custom_errors"
	"russian-roulette/internal/entities/types"
	userEntities "russian-roulette/internal/entities/user"
)

// todo сделать
func (cmd *Commands) Pull(ctx context.Context, message *tgbotapi.Message) error {
	u, ok := ctx.Value(middleware.UserContextKey).(*userEntities.User)
	if !ok {
		return errors.New(projectCustomErrors.ErrUserNotFoundInContext)
	}

	game, err := cmd.userService.GetUserActiveGame(ctx, u.Uuid)
	if err != nil {
		if err.Error() == projectCustomErrors.ErrGameNotFound {
			return custom_errors.ErrUserWithoutActiveGame
		}
		return err
	}

	msgText := ""

	// todo add error messages
	isDead, isOver, currentPlayer, nextPlayer, game, err := cmd.gameService.PullTrigger(ctx, game, u.Uuid)
	if err != nil {
		return err
	}
	cmd.logger.Debug("log all output pull trigger", zap.Bool("isDead", isDead), zap.Bool("isOver", isOver))
	if isOver {
		players := []string{currentPlayer.Name, nextPlayer.Name}
		msgText = text.FinishGameMessage(nextPlayer.Name, players, types.DecimalToInt(&game.BetAmount), game.BulletCount)
	}
	if isDead {
		msgText = text.UnsuccessfulPullMessage(currentPlayer.Name, nextPlayer.Name, game.BulletCount, game.RoundsCount)
	}
	if !isOver && !isDead {
		msgText = text.SuccessfulPullMessage(currentPlayer.Name, game.BulletCount, game.RoundsCount)
	}

	err = telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
		ChatId: message.Chat.ID,
		Text:   msgText,
	}, cmd.logger)
	if err != nil {
		return err
	}

	return nil
}

// todo сделать
func (cmd *Commands) Pass(ctx context.Context, message *tgbotapi.Message) error {
	u, ok := ctx.Value(middleware.UserContextKey).(*userEntities.User)
	if !ok {
		return errors.New(projectCustomErrors.ErrUserNotFoundInContext)
	}

	game, err := cmd.userService.GetUserActiveGame(ctx, u.Uuid)
	if err != nil {
		if err.Error() == projectCustomErrors.ErrGameNotFound {
			return custom_errors.ErrUserWithoutActiveGame
		}
		return err
	}

	currentPlayer, nextPlayer, bullets, rounds, err := cmd.gameService.PassTrigger(ctx, game.Uuid, u.Uuid)
	if err != nil {
		return err
	}

	// todo send message to this player and next
	err = telegramUtils.SendMessage(cmd.bot, &telegramUtils.Message{
		ChatId: message.Chat.ID,
		Text:   text.SuccessfulPassMessage(currentPlayer.Name, nextPlayer.Name, bullets, rounds),
	}, cmd.logger)
	if err != nil {
		return err
	}

	return nil
}
