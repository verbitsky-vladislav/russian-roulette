package game

import "time"

type Game struct {
	ID          string    `json:"id"`
	CreatorID   string    `json:"creator_id"`
	Status      string    `json:"status"`
	BetAmount   float64   `json:"bet_amount"`
	BulletCount int       `json:"bullet_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateGame struct {
	CreatorID   string  `json:"creator_id"`
	Status      string  `json:"status"`
	BetAmount   float64 `json:"bet_amount"`
	BulletCount int     `json:"bullet_count"`
}

type UpdateGame struct {
	Status      *string  `json:"status,omitempty"`
	BetAmount   *float64 `json:"bet_amount,omitempty"`
	BulletCount *int     `json:"bullet_count,omitempty"`
}

type GetGameFilters struct {
	CreatorID *string `json:"creator_id,omitempty"`
	Status    *string `json:"status,omitempty"`
}
