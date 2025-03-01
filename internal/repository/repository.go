package repository

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"russian-roulette/internal/config"
	userEntities "russian-roulette/internal/entities/user"
)

type database struct {
	DB     *sql.DB
	logger *zap.Logger
}

func New(cfg *config.Database, logger *zap.Logger) *database {
	db, err := sql.Open("postgres", cfg.DNS)
	if err != nil {
		logger.Panic("failed to connect to database", zap.Error(err))
	}

	err = db.Ping()
	if err != nil {
		logger.Panic("failed to ping to database", zap.Error(err))
	}

	return &database{
		DB:     db,
		logger: logger,
	}
}

func (d *database) Close() error {
	return d.DB.Close()
}

type (
	UserRepository interface {
		Create(ctx context.Context, newUser *userEntities.CreateUser) (*userEntities.User, error)
		Update(ctx context.Context, updUser userEntities.UpdateUser) (*userEntities.User, error)
		GetByUUID(ctx context.Context, uuid string) (*userEntities.User, error)
		GetByChatId(ctx context.Context, chatId int64) (*userEntities.User, error)
		GetAll(ctx context.Context, filters *userEntities.GetUserFilters) ([]*userEntities.User, error)
	}
)
