package game

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	dec "github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
	"russian-roulette/internal/entities/custom_errors"
	gameEntities "russian-roulette/internal/entities/game"
	"russian-roulette/internal/entities/types"
	"russian-roulette/internal/models"
)

type GameRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewGameRepository(db *sql.DB, logger *zap.Logger) *GameRepository {
	return &GameRepository{
		db:     db,
		logger: logger,
	}
}

func (r *GameRepository) NewFromModel(model *models.Game) (*gameEntities.Game, error) {

	return &gameEntities.Game{
		Uuid:        model.UUID,
		CreatorUuid: model.CreatorUUID.String,
		Status:      gameEntities.GameStatus(model.Status),
		BetAmount:   *types.NewDecimalFromString(model.BetAmount.String()),
		BulletCount: model.BulletCount,
		RoundsCount: model.RoundsCount,
		CreatedAt:   model.CreatedAt.Time,
	}, nil
}

func (r *GameRepository) Create(ctx context.Context, newGame *gameEntities.CreateGame) (*gameEntities.Game, error) {
	r.logger.Debug("create new game", zap.Any("game", newGame))

	bet, err := dec.NewFromString(newGame.BetAmount.String())
	if err != nil {
		return nil, err
	}
	game := &models.Game{
		CreatorUUID: null.NewString(newGame.CreatorUuid, newGame.CreatorUuid != ""),
		Status:      string(newGame.Status),
		BetAmount:   bet,
		BulletCount: newGame.BulletCount,
		RoundsCount: newGame.RoundsCount,
	}

	err = game.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrInsertGame)
	}

	result, err := r.NewFromModel(game)
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrInsertGame)
	}

	return result, nil

}

func (r *GameRepository) Update(ctx context.Context, upd *gameEntities.UpdateGame) (*gameEntities.Game, error) {
	game, err := models.Games(models.GameWhere.UUID.EQ(upd.Uuid)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(custom_errors.ErrGameNotFound)
		}
		return nil, errors.Wrap(err, custom_errors.ErrUpdateGame)
	}

	if upd.Status != nil {
		game.Status = string(*upd.Status)
	}
	if upd.BetAmount != nil {
		bet, err := dec.NewFromString(upd.BetAmount.String())
		if err != nil {
			return nil, err
		}
		game.BetAmount = bet
	}
	if upd.BulletCount != nil {
		game.BulletCount = *upd.BulletCount
	}
	if upd.RoundsCount != nil {
		game.RoundsCount = *upd.RoundsCount
	}

	_, err = game.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrUpdateGame)
	}

	result, err := r.NewFromModel(game)
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrUpdateGame)
	}

	return result, nil
}

func (r *GameRepository) GetAll(ctx context.Context, filters *gameEntities.GetGameFilters) ([]*gameEntities.Game, error) {
	var qms []qm.QueryMod

	if filters.Uuid != nil {
		qms = append(qms, models.GameWhere.UUID.EQ(*filters.Uuid))
	}

	if filters.CreatorUuid != nil {
		qms = append(qms, models.GameWhere.CreatorUUID.EQ(null.NewString(*filters.CreatorUuid, *filters.CreatorUuid != "")))
	}

	if filters.Status != nil {
		qms = append(qms, models.GameWhere.Status.EQ(*filters.Status))
	}

	if filters.BetAmount != nil {
		bet, err := dec.NewFromString(filters.BetAmount.String())
		if err != nil {
			return nil, err
		}
		qms = append(qms, models.GameWhere.BetAmount.EQ(bet))
	}

	if filters.BulletCount != nil {
		qms = append(qms, models.GameWhere.BulletCount.EQ(*filters.BulletCount))
	}

	if filters.RoundsCount != nil {
		qms = append(qms, models.GameWhere.RoundsCount.EQ(*filters.RoundsCount))
	}

	if filters.Limit != 0 {
		qms = append(qms, qm.Limit(filters.Limit))
	}

	if filters.Offset != 0 {
		qms = append(qms, qm.Offset(filters.Offset))
	}

	// Если указан userUuid — добавляем INNER JOIN + фильтрацию по игрокам
	if filters.UserUuid != nil {
		qms = append(qms,
			qm.InnerJoin("game_players ON game_players.game_uuid = game.uuid"),
			qm.Where("game_players.user_uuid = ?", *filters.UserUuid),
		)
	}

	games, err := models.Games(qms...).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, custom_errors.ErrFetchGames)
	}

	result := make([]*gameEntities.Game, 0, len(games))

	for _, game := range games {
		u, err := r.NewFromModel(game)
		if err != nil {
			return nil, errors.Wrap(err, custom_errors.ErrFetchGames)
		}
		result = append(result, u)
	}

	return result, nil
}

func (r *GameRepository) GetByUUID(ctx context.Context, uuid string) (*gameEntities.Game, error) {
	users, err := r.GetAll(ctx, &gameEntities.GetGameFilters{Uuid: &uuid})
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrGetGameByUuid)
	}

	if len(users) == 0 {
		return nil, errors.New(custom_errors.ErrGameNotFound)
	}

	return users[0], nil
}
