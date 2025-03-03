package game

import (
	"russian-roulette/internal/entities/types"
	"time"
)

type GameStatus string

var (
	Waiting  GameStatus = "waiting"
	Active   GameStatus = "active"
	Finished GameStatus = "finished"
	Canceled GameStatus = "cancelled"
)

type Game struct {
	Uuid        string        `json:"uuid"`
	CreatorUuid string        `json:"creator_uuid"`
	Status      GameStatus    `json:"status"`
	BetAmount   types.Decimal `json:"bet_amount"`
	BulletCount int           `json:"bullet_count"`
	CreatedAt   time.Time     `json:"created_at"`
}

type CreateGame struct {
	CreatorUuid string        `json:"creator_uuid"`
	Status      GameStatus    `json:"status"`
	BetAmount   types.Decimal `json:"bet_amount"`
	BulletCount int           `json:"bullet_count"`
}

type UpdateGame struct {
	Uuid        string         `json:"uuid"`
	Status      *GameStatus    `json:"status,omitempty"`
	BetAmount   *types.Decimal `json:"bet_amount,omitempty"`
	BulletCount *int           `json:"bullet_count,omitempty"`
}

type GetGameFilters struct {
	Uuid          *string        `query:"uuid"`
	CreatorUuid   *string        `json:"creator_uuid,omitempty"`
	Status        *string        `json:"status,omitempty"`
	BetAmount     *types.Decimal `json:"bet_amount,omitempty"`
	BulletCount   *int           `json:"bullet_count,omitempty"`
	CreatedAtFrom *time.Time     `json:"created_at_from,omitempty"`
	CreatedAtTo   *time.Time     `json:"created_at_to,omitempty"`
	Limit         int            `json:"limit"`
	Offset        int            `json:"offset"`
}
