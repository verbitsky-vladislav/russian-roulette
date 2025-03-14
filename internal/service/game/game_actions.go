package game

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"go.uber.org/zap"
	"math/rand"
	gameEntities "russian-roulette/internal/entities/game"
	projectUtils "russian-roulette/internal/utils"
)

func (g *GameService) AddUserToGame(ctx context.Context, userUuid, gameUuid, name string) (*gameEntities.GamePlayer, error) {
	player, err := g.gamePlayerRepository.Create(ctx, &gameEntities.CreateGamePlayer{
		Name:     name,
		GameUuid: gameUuid,
		UserUuid: userUuid,
	})
	if err != nil {
		return nil, err
	}

	return player, nil
}

// todo объединить в транзакцию вместе с redis
func (g *GameService) PullTrigger(
	ctx context.Context,
	game *gameEntities.Game,
	playerUuid string,
) (isDead, isOver bool, currentPlayer, nextPlayer *gameEntities.GamePlayer, updatedGame *gameEntities.Game, err error) {

	g.logger.Info("PullTrigger called", zap.String("game_uuid", game.Uuid), zap.String("player_uuid", playerUuid))

	if game.Uuid == "" {
		err = errors.New("game UUID is empty")
		g.logger.Error("Invalid game UUID", zap.Error(err))
		return false, false, nil, nil, nil, err
	}

	// Получаем игроков
	g.logger.Debug("Fetching current and next players", zap.String("game_uuid", game.Uuid))
	currentPlayer, nextPlayer, err = g.getCurrentAndNextPlayers(ctx, game.Uuid)
	if err != nil {
		g.logger.Error("Failed to get players", zap.String("game_uuid", game.Uuid), zap.Error(err))
		return false, false, nil, nil, nil, fmt.Errorf("failed to get players: %w", err)
	}
	g.logger.Debug("Players fetched successfully", zap.Any("current_player", currentPlayer), zap.Any("next_player", nextPlayer))

	if currentPlayer == nil || nextPlayer == nil {
		err = errors.New("failed to retrieve current or next player")
		g.logger.Error("Invalid player data", zap.Any("current", currentPlayer), zap.Any("next", nextPlayer))
		return false, false, currentPlayer, nextPlayer, nil, err
	}

	// Получаем текущего игрока
	g.logger.Debug("Fetching current player's turn", zap.String("game_uuid", game.Uuid))
	currentPlayerUuid, err := g.getCurrentTurn(game.Uuid)
	if err != nil {
		g.logger.Error("Failed to get current turn", zap.String("game_uuid", game.Uuid), zap.Error(err))
		return false, false, nil, nil, nil, fmt.Errorf("failed to get current turn: %w", err)
	}
	g.logger.Debug("Current player's turn fetched", zap.String("current_player_uuid", currentPlayerUuid))

	// Проверяем, что запрос отправил именно текущий игрок
	if playerUuid != currentPlayerUuid {
		err = errors.New("action not allowed: not player's turn")
		g.logger.Warn("Player attempted action out of turn", zap.String("expected", currentPlayerUuid), zap.String("actual", playerUuid))
		return false, false, nil, nil, nil, err
	}

	// Проверка количества патронов
	if game.BulletCount <= 0 {
		err = errors.New("invalid bullet count (<= 0)")
		g.logger.Error("Invalid bullet count", zap.Int("bullets", game.BulletCount))
		return false, false, nil, nil, nil, err
	}

	isDead = false
	shotResult := gameEntities.Miss

	// Выстрел: 1 из оставшихся патронов
	g.logger.Debug("Player pulling trigger", zap.String("player_uuid", currentPlayerUuid))
	if rand.Intn(game.RoundsCount) < game.BulletCount {
		g.logger.Info("Player shot and died", zap.String("player_uuid", currentPlayerUuid))

		// Удаляем из очереди
		if err = g.removeFromQueue(game.Uuid, currentPlayerUuid); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				g.logger.Warn("Player not found in queue", zap.String("player_uuid", currentPlayerUuid))
			} else {
				g.logger.Error("Failed to remove player from queue", zap.Error(err))
				return false, false, nil, nil, nil, fmt.Errorf("failed to remove player from queue: %w", err)
			}
		}

		// Фиксируем смерть игрока
		_, err = g.gamePlayerRepository.Update(ctx, &gameEntities.UpdateGamePlayer{
			UserUuid: currentPlayerUuid,
			HasShot:  true,
			IsAlive:  false,
		})
		if err != nil {
			g.logger.Error("Failed to update player status", zap.String("player_uuid", currentPlayerUuid), zap.Error(err))
			return false, false, nil, nil, nil, fmt.Errorf("failed to update player status: %w", err)
		}

		isDead = true
		shotResult = gameEntities.Shot
	}

	// Проверяем, остался ли один игрок (победитель)
	if isDead {
		g.logger.Debug("Checking if game is over", zap.String("game_uuid", game.Uuid))
		isOver, _, err = g.isGameOver(game.Uuid)
		if err != nil {
			g.logger.Error("Failed to check game over", zap.String("game_uuid", game.Uuid), zap.Error(err))
			return false, false, nil, nil, nil, fmt.Errorf("failed to check game over: %w", err)
		}

		if isOver {
			g.logger.Info("Game over", zap.String("winner_uuid", nextPlayer.UserUuid))
			if err = g.FinishGame(ctx, game.Uuid, nextPlayer.UserUuid); err != nil {
				g.logger.Error("Failed to finish game", zap.String("game_uuid", game.Uuid), zap.Error(err))
				return false, false, nil, nil, nil, fmt.Errorf("failed to finish game: %w", err)
			}
			return isDead, isOver, currentPlayer, nextPlayer, nil, nil
		}
	}

	// Обновляем количество патронов
	updGame := &gameEntities.UpdateGame{
		Uuid:        game.Uuid,
		RoundsCount: projectUtils.ToPtr(game.RoundsCount - 1),
	}

	if isDead {
		updGame.BulletCount = projectUtils.ToPtr(game.BulletCount - 1)
	}

	// Создаем раунд
	g.logger.Debug("Creating game round", zap.String("game_uuid", game.Uuid))
	_, err = g.CreateRound(ctx, &gameEntities.CreateGameRound{
		GameUuid: game.Uuid,
		UserUuid: currentPlayerUuid,
		Action:   projectUtils.ToPtr(gameEntities.Pull),
		Result:   &shotResult,
	})
	if err != nil {
		g.logger.Error("Failed to create round", zap.String("game_uuid", game.Uuid), zap.Error(err))
		return false, false, nil, nil, nil, fmt.Errorf("failed to create round: %w", err)
	}

	// Обновляем игру
	g.logger.Debug("Updating game state", zap.String("game_uuid", game.Uuid))
	game, err = g.gameRepository.Update(ctx, updGame)
	if err != nil {
		g.logger.Error("Failed to update game", zap.String("game_uuid", game.Uuid), zap.Error(err))
		return false, false, nil, nil, nil, fmt.Errorf("failed to update game: %w", err)
	}

	g.logger.Info("PullTrigger completed successfully", zap.String("game_uuid", game.Uuid), zap.String("player_uuid", currentPlayerUuid), zap.Bool("isDead", isDead), zap.Bool("isOver", isOver))

	return isDead, isOver, currentPlayer, nextPlayer, game, nil
}

// todo объединить в транзакцию вместе с редисом
func (g *GameService) PassTrigger(ctx context.Context, gameUuid, userUuid string) (currentPlayer, nextPlayer *gameEntities.GamePlayer, bullets, rounds int, err error) {
	// Получаем последний раунд игры
	lastRound, err := g.GetLastRound(ctx, gameUuid)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	// Проверяем, что последний ход сделал этот же игрок и он не пасовал ранее
	if lastRound != nil && lastRound.UserUuid == userUuid && lastRound.Action == gameEntities.Pass {
		return nil, nil, 0, 0, errors.New("нельзя дважды подряд пасовать")
	}

	game, _, _, err := g.GetGameByUuid(ctx, gameUuid, false, false)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	currentPlayer, nextPlayer, err = g.getCurrentAndNextPlayers(ctx, gameUuid)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	// Передаем ход следующему игроку
	nextPlayerUuid, err := g.getNextTurn(gameUuid)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	if nextPlayer.UserUuid != nextPlayerUuid {
		g.logger.Debug("debug:", zap.Any("currentPlayer:", currentPlayer), zap.Any("nextPlayer:", nextPlayer), zap.String("nexrplayerguid", nextPlayerUuid))
		return nil, nil, 0, 0, errors.New("в общем какая-то лажа с current / next пользователями -> лажа с очередями - смотри и изучай")
	}

	// Создаем новый раунд с действием "пас"
	_, err = g.CreateRound(ctx, &gameEntities.CreateGameRound{
		GameUuid: gameUuid,
		UserUuid: userUuid,
		Action:   projectUtils.ToPtr(gameEntities.Pass),
	})

	return currentPlayer, nextPlayer, game.BulletCount, game.RoundsCount, err
}

func (g *GameService) createTurnsQueue(gameUuid string, players []gameEntities.GamePlayer) error {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	for _, p := range players {
		err := g.cacheService.PushToQueue(redisKey, p.UserUuid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GameService) getCurrentTurn(gameUuid string) (string, error) {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	currentPlayerUuid, err := g.cacheService.PeekQueue(redisKey)
	if err != nil {
		return "", err
	}

	if currentPlayerUuid == "" {
		return "", errors.New("очередь ходов пуста")
	}

	err = g.cacheService.PushToQueue(redisKey, currentPlayerUuid)
	if err != nil {
		return "", err
	}

	return currentPlayerUuid, nil
}

func (g *GameService) getNextTurn(gameUuid string) (string, error) {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	_, err := g.cacheService.PopFromQueue(redisKey)
	if err != nil {
		return "", err
	}

	currentPlayer, err := g.cacheService.PeekQueue(redisKey)
	if err != nil {
		return "", err
	}

	return currentPlayer, nil
}

func (g *GameService) removeFromQueue(gameUuid, userUuid string) error {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	return g.cacheService.RemoveFromQueue(redisKey, userUuid)
}

// todo не присылается повесть о победе -> правильно надо валидировать
func (g *GameService) isGameOver(gameUuid string) (bool, string, error) {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	players, err := g.cacheService.GetQueueValues(redisKey)
	if err != nil {
		return false, "", err
	}

	if len(players) == 1 {
		return true, players[0], nil
	}

	return false, "", nil
}

func (g *GameService) getCurrentAndNextPlayers(ctx context.Context, gameUuid string) (current, next *gameEntities.GamePlayer, err error) {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	players, err := g.cacheService.GetQueueValues(redisKey)
	if err != nil {
		return nil, nil, err
	}

	if len(players) == 0 {
		return nil, nil, errors.New("очередь пуста")
	}

	currentUuid := players[0]
	var nextUuid string
	if len(players) > 1 {
		nextUuid = players[1]
	} else {
		nextUuid = players[0] // Если остался один игрок, он и будет следующим (но в целом игра должна завершиться)
	}

	// Получаем объекты игроков
	currentPlayer, err := g.gamePlayerRepository.GetAll(ctx, &gameEntities.GetGamePlayersFilters{
		GameUuid: projectUtils.ToPtr(gameUuid),
		UserUuid: projectUtils.ToPtr(currentUuid),
	})
	if err != nil {
		return nil, nil, err
	}

	nextPlayer, err := g.gamePlayerRepository.GetAll(ctx, &gameEntities.GetGamePlayersFilters{
		GameUuid: projectUtils.ToPtr(gameUuid),
		UserUuid: projectUtils.ToPtr(nextUuid),
	})
	if err != nil {
		return nil, nil, err
	}

	return currentPlayer[0], nextPlayer[0], nil
}
