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
	"russian-roulette/internal/entities/game"
	"russian-roulette/internal/models"
)

type GamePlayersRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewGamePlayersRepository(db *sql.DB, logger *zap.Logger) *GamePlayersRepository {
	return &GamePlayersRepository{
		db:     db,
		logger: logger,
	}
}

func (r *GamePlayersRepository) NewFromModel(model *models.GamePlayer) (*game.GamePlayer, error) {
	return &game.GamePlayer{
		Uuid:     model.UUID,
		GameUuid: model.GameUUID.String,
		UserUuid: model.UserUUID.String,
		HasShot:  model.HasShot.Bool,
		IsAlive:  model.IsAlive.Bool,
		Name:     model.Name,
	}, nil
}

func (r *GamePlayersRepository) Create(ctx context.Context, newPlayer *game.CreateGamePlayer) (*game.GamePlayer, error) {
	player := &models.GamePlayer{
		GameUUID: null.NewString(newPlayer.GameUuid, newPlayer.GameUuid != ""),
		UserUUID: null.NewString(newPlayer.UserUuid, newPlayer.UserUuid != ""),
		Name:     newPlayer.Name,
		IsAlive:  null.NewBool(true, true),
	}

	err := player.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrInsertGamePlayers)
	}

	return r.NewFromModel(player)
}

func (r *GamePlayersRepository) Update(ctx context.Context, upd *game.UpdateGamePlayer) (*game.GamePlayer, error) {
	player, err := models.GamePlayers(models.GamePlayerWhere.UserUUID.EQ(null.NewString(upd.UserUuid, upd.UserUuid != ""))).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(custom_errors.ErrGamePlayersNotFound)
		}
		return nil, errors.Wrap(err, custom_errors.ErrUpdateGamePlayers)
	}

	player.HasShot = null.NewBool(upd.HasShot, true)
	player.IsAlive = null.NewBool(upd.IsAlive, true)

	_, err = player.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrUpdateGamePlayers)
	}

	return r.NewFromModel(player)
}

func (r *GamePlayersRepository) GetAll(ctx context.Context, filters *game.GetGamePlayersFilters) ([]*game.GamePlayer, error) {
	var qms []qm.QueryMod

	if filters.Uuid != nil {
		qms = append(qms, models.GamePlayerWhere.UUID.EQ(*filters.Uuid))
	}
	if filters.GameUuid != nil {
		qms = append(qms, models.GamePlayerWhere.GameUUID.EQ(null.NewString(*filters.GameUuid, *filters.GameUuid != "")))
	}
	if filters.UserUuid != nil {
		qms = append(qms, models.GamePlayerWhere.UserUUID.EQ(null.NewString(*filters.UserUuid, *filters.UserUuid != "")))
	}
	if filters.Name != nil {
		qms = append(qms, models.GamePlayerWhere.Name.EQ(*filters.Name))
	}
	if filters.HasShot != nil {
		qms = append(qms, models.GamePlayerWhere.HasShot.EQ(null.NewBool(*filters.HasShot, filters.HasShot != nil)))
	}
	if filters.IsAlive != nil {
		qms = append(qms, models.GamePlayerWhere.IsAlive.EQ(null.NewBool(*filters.IsAlive, filters.IsAlive != nil)))
	}

	players, err := models.GamePlayers(qms...).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, custom_errors.ErrFetchGamePlayers)
	}

	result := make([]*game.GamePlayer, 0, len(players))
	for _, player := range players {
		p, err := r.NewFromModel(player)
		if err != nil {
			return nil, errors.Wrap(err, custom_errors.ErrFetchGamePlayers)
		}
		result = append(result, p)
	}

	return result, nil
}

func (r *GamePlayersRepository) GetByGameUUID(ctx context.Context, gameUuid string) ([]*game.GamePlayer, error) {
	return r.GetAll(ctx, &game.GetGamePlayersFilters{GameUuid: &gameUuid})
}
