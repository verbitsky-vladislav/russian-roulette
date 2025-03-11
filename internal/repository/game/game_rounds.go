package game

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
	"russian-roulette/internal/entities/custom_errors"
	gameEntities "russian-roulette/internal/entities/game"
	"russian-roulette/internal/models"
)

type GameRoundRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewGameRoundRepository(db *sql.DB, logger *zap.Logger) *GameRoundRepository {
	return &GameRoundRepository{db: db, logger: logger}
}

func (r *GameRoundRepository) NewFromModel(model *models.GameRound) (*gameEntities.GameRound, error) {
	return &gameEntities.GameRound{
		Uuid:     model.UUID,
		GameUuid: model.GameUUID.String,
		UserUuid: model.UserUUID.String,
		Action:   gameEntities.GameAction(model.Action.String),
		Result:   gameEntities.GameActionResult(model.Result.String),
		Created:  model.CreatedAt.Time,
	}, nil
}

func (r *GameRoundRepository) Create(ctx context.Context, newRound *gameEntities.CreateGameRound) (*gameEntities.GameRound, error) {
	round := &models.GameRound{
		GameUUID: null.NewString(newRound.GameUuid, newRound.GameUuid != ""),
		UserUUID: null.NewString(newRound.UserUuid, newRound.UserUuid != ""),
	}

	err := round.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrInsertGameRounds)
	}

	return r.NewFromModel(round)
}

func (r *GameRoundRepository) Update(ctx context.Context, upd *gameEntities.UpdateGameRound) (*gameEntities.GameRound, error) {
	round, err := models.GameRounds(models.GameRoundWhere.UUID.EQ(upd.Uuid)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(custom_errors.ErrGameRoundsNotFound)
		}
		return nil, errors.Wrap(err, custom_errors.ErrUpdateGameRounds)
	}

	if upd.Action != nil {
		round.Action = null.NewString(string(*upd.Action), *upd.Action != "")
	}

	if upd.Result != nil {
		round.Result = null.NewString(string(*upd.Result), *upd.Result != "")
	}

	_, err = round.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrUpdateGameRounds)
	}

	return r.NewFromModel(round)
}

func (r *GameRoundRepository) GetAll(ctx context.Context, filters *gameEntities.GetGameRoundsFilters) ([]*gameEntities.GameRound, error) {
	var qms []qm.QueryMod

	if filters.Uuid != nil {
		qms = append(qms, models.GameRoundWhere.UUID.EQ(*filters.Uuid))
	}
	if filters.GameUuid != nil {
		qms = append(qms, models.GameRoundWhere.GameUUID.EQ(null.NewString(*filters.GameUuid, *filters.GameUuid != "")))
	}
	if filters.UserUuid != nil {
		qms = append(qms, models.GameRoundWhere.UserUUID.EQ(null.NewString(*filters.Uuid, *filters.Uuid != "")))
	}
	if filters.Action != nil {
		qms = append(qms, models.GameRoundWhere.Action.EQ(null.NewString(string(*filters.Action), *filters.Action != "")))
	}
	if filters.Result != nil {
		qms = append(qms, models.GameRoundWhere.Result.EQ(null.NewString(string(*filters.Result), *filters.Result != "")))
	}

	rounds, err := models.GameRounds(qms...).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, custom_errors.ErrFetchGamesRounds)
	}

	result := make([]*gameEntities.GameRound, 0, len(rounds))
	for _, round := range rounds {
		r, err := r.NewFromModel(round)
		if err != nil {
			return nil, errors.Wrap(err, custom_errors.ErrFetchGamesRounds)
		}
		result = append(result, r)
	}

	return result, nil
}
