package user

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	boilerTypes "github.com/volatiletech/sqlboiler/v4/types"
	"go.uber.org/zap"
	"russian-roulette/internal/entities/custom_errors"
	"russian-roulette/internal/entities/types"
	userEntities "russian-roulette/internal/entities/user"
	"russian-roulette/internal/models"
)

type UserRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewUserRepository(db *sql.DB, logger *zap.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) NewFromModel(model *models.User) (*userEntities.User, error) {
	return &userEntities.User{
		Uuid:          model.UUID,
		ChatId:        int64(model.ChatID),
		TgName:        model.TGName,
		Balance:       types.Decimal{Big: model.Balance.Big},
		WalletAddress: model.WalletAddress.String,
		TotalWins:     int64(model.TotalWins.Int),
		TotalLosses:   int64(model.TotalLosses.Int),
		CreatedAt:     model.CreatedAt.Time,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, newUser *userEntities.CreateUser) (*userEntities.User, error) {
	user := &models.User{
		ChatID: int(newUser.ChatId),
		TGName: newUser.TgName,
	}

	err := user.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrInsertUser)
	}

	return r.NewFromModel(user)
}

func (r *UserRepository) Update(ctx context.Context, upd *userEntities.UpdateUser) (*userEntities.User, error) {
	user, err := models.Users(models.UserWhere.ChatID.EQ(int(upd.ChatId))).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(custom_errors.ErrUserNotFound)
		}
		return nil, errors.Wrap(err, custom_errors.ErrUpdateUser)
	}

	if upd.TgName != nil {
		user.TGName = *upd.TgName
	}
	if upd.Balance != nil {
		user.Balance = boilerTypes.NullDecimal(*upd.Balance)
	}
	if upd.WalletAddress != nil {
		user.WalletAddress = null.NewString(*upd.WalletAddress, *upd.WalletAddress != "")
	}
	if upd.TotalWins != nil {
		user.TotalWins = null.NewInt(int(*upd.TotalWins), *upd.TotalWins > 0)
	}
	if upd.TotalLosses != nil {
		user.TotalLosses = null.NewInt(int(*upd.TotalLosses), *upd.TotalLosses > 0)
	}

	_, err = user.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrUpdateUser)
	}

	return r.NewFromModel(user)
}

func (r *UserRepository) GetAll(ctx context.Context, filters *userEntities.GetUserFilters) ([]*userEntities.User, error) {
	var qms []qm.QueryMod

	if filters.Uuid != nil {
		qms = append(qms, models.UserWhere.UUID.EQ(*filters.Uuid))
	}
	if filters.ChatId != nil {
		qms = append(qms, models.UserWhere.ChatID.EQ(int(*filters.ChatId)))
	}
	if filters.TgName != nil {
		qms = append(qms, models.UserWhere.TGName.EQ(*filters.TgName))
	}
	if filters.Balance != nil {
		qms = append(qms, models.UserWhere.Balance.EQ(boilerTypes.NullDecimal(boilerTypes.NewDecimal(filters.Balance.Big))))
	}
	if filters.WalletAddress != nil {
		qms = append(qms, models.UserWhere.WalletAddress.EQ(null.NewString(*filters.WalletAddress, *filters.WalletAddress != "")))
	}
	if filters.TotalWins != nil {
		qms = append(qms, models.UserWhere.TotalWins.EQ(null.NewInt(int(*filters.TotalWins), *filters.TotalWins > 0)))
	}
	if filters.TotalLosses != nil {
		qms = append(qms, models.UserWhere.TotalLosses.EQ(null.NewInt(int(*filters.TotalLosses), *filters.TotalLosses > 0)))
	}
	if filters.Limit != nil {
		qms = append(qms, qm.Limit(*filters.Limit))
	}
	if filters.Offset != nil {
		qms = append(qms, qm.Offset(*filters.Offset))
	}

	users, err := models.Users(qms...).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, custom_errors.ErrFetchUsers)
	}

	result := make([]*userEntities.User, 0, len(users))

	for _, user := range users {
		u, err := r.NewFromModel(user)
		if err != nil {
			return nil, errors.Wrap(err, custom_errors.ErrFetchUsers)
		}
		result = append(result, u)
	}

	return result, nil
}

func (r *UserRepository) GetByUUID(ctx context.Context, uuid string) (*userEntities.User, error) {
	users, err := r.GetAll(ctx, &userEntities.GetUserFilters{Uuid: &uuid})
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrFetchUsers)
	}

	if len(users) == 0 {
		return nil, errors.New(custom_errors.ErrUserNotFound)
	}

	return users[0], nil
}

func (r *UserRepository) GetByChatID(ctx context.Context, chatId int64) (*userEntities.User, error) {
	users, err := r.GetAll(ctx, &userEntities.GetUserFilters{ChatId: &chatId})
	if err != nil {
		return nil, errors.Wrap(err, custom_errors.ErrFetchUsers)
	}

	if len(users) == 0 {
		return nil, errors.New(custom_errors.ErrUserNotFound)
	}

	return users[0], nil
}
