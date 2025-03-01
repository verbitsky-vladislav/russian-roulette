package game

import (
	"database/sql"
	"go.uber.org/zap"
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
