package game

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
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

// todo add error handling
func (g *GameService) PullTrigger(ctx context.Context, game *gameEntities.Game) (isDead, isOver bool, err error) {
	isDead = false
	isOver = false
	shotResult := gameEntities.Miss

	// Получаем текущего игрока
	currentPlayerUuid, err := g.getNextTurn(game.Uuid)
	if err != nil {
		return isDead, isOver, err
	}

	// Определяем, произошел ли выстрел (1 шанс из оставшихся патронов)
	remainingBullets := game.BulletCount // В идеале это должно храниться в БД или кэше
	if rand.Intn(6) < remainingBullets { // Например, 1 из 6
		// Игрок проиграл
		// удаляем из очереди
		err = g.removeFromQueue(game.Uuid, currentPlayerUuid)
		if err != nil {
			return isDead, isOver, err
		}
		// фиксируем его смерть
		_, err = g.gamePlayerRepository.Update(ctx, &gameEntities.UpdateGamePlayer{
			UserUuid: currentPlayerUuid,
			HasShot:  true,
			IsAlive:  false,
		})

		// Проверяем, остался ли один игрок (победитель)
		isOver, _, err = g.isGameOver(game.Uuid)
		if err != nil {
			return isDead, isOver, err
		}

		if isOver { // func finish game
			err = g.FinishGame(ctx, currentPlayerUuid)
			if err != nil {
				return isDead, isOver, err
			}
			return isDead, isOver, err
		}

		isDead = true
		shotResult = gameEntities.Shot
	}

	// Создаем раунд
	_, err = g.CreateRound(ctx, &gameEntities.CreateGameRound{
		GameUuid: game.Uuid,
		UserUuid: currentPlayerUuid,
		Action:   projectUtils.ToPtr(gameEntities.Pull),
		Result:   &shotResult,
	})

	return isDead, isOver, err
}

func (g *GameService) PassTrigger(ctx context.Context, gameUuid, userUuid string) (nextPlayer *gameEntities.GamePlayer, err error) {
	// Получаем последний раунд игры
	lastRound, err := g.GetLastRound(ctx, gameUuid)
	if err != nil {
		return nil, err
	}

	// Проверяем, что последний ход сделал этот же игрок и он не пасовал ранее
	if lastRound != nil && lastRound.UserUuid == userUuid && lastRound.Action == gameEntities.Pass {
		return nil, errors.New("нельзя дважды подряд пасовать")
	}

	// Передаем ход следующему игроку
	nextPlayerUuid, err := g.getNextTurn(gameUuid)
	if err != nil {
		return nil, err
	}

	player, err := g.gamePlayerRepository.GetAll(ctx, &gameEntities.GetGamePlayersFilters{
		GameUuid: projectUtils.ToPtr(gameUuid),
		UserUuid: projectUtils.ToPtr(nextPlayerUuid),
	})

	// Создаем новый раунд с действием "пас"
	_, err = g.CreateRound(ctx, &gameEntities.CreateGameRound{
		GameUuid: gameUuid,
		UserUuid: userUuid,
		Action:   projectUtils.ToPtr(gameEntities.Pass),
	})

	return player[0], err
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

func (g *GameService) getNextTurn(gameUuid string) (string, error) {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	nextPlayerUuid, err := g.cacheService.PopFromQueue(redisKey)
	if err != nil {
		return "", err
	}

	if nextPlayerUuid == "" {
		return "", errors.New("очередь ходов пуста")
	}

	err = g.cacheService.PushToQueue(redisKey, nextPlayerUuid)
	if err != nil {
		return "", err
	}

	return nextPlayerUuid, nil
}

func (g *GameService) removeFromQueue(gameUuid, userUuid string) error {
	redisKey := fmt.Sprintf("game:%s:turns", gameUuid)

	return g.cacheService.RemoveFromQueue(redisKey, userUuid)
}

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
