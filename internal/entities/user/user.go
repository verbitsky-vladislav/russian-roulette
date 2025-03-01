package user

import (
	"russian-roulette/internal/entities/types"
	"time"
)

type User struct {
	Uuid          string        `json:"uuid"`
	ChatId        int64         `json:"chat_id"`
	TgName        string        `json:"tg_name"`
	Balance       types.Decimal `json:"balance"`
	WalletAddress string        `json:"wallet_address"`
	TotalWins     int64         `json:"total_wins"`
	TotalLosses   int64         `json:"total_losses"`
	CreatedAt     time.Time     `json:"created_at"`
}

type CreateUser struct {
	ChatId int64  `json:"chat_id"`
	TgName string `json:"tg_name"`
}

type UpdateUser struct {
	ChatId        int64          `json:"chat_id"`
	TgName        *string        `json:"tg_name"`
	Balance       *types.Decimal `json:"balance"`
	WalletAddress *string        `json:"wallet_address"`
	TotalWins     *int64         `json:"total_wins"`
	TotalLosses   *int64         `json:"total_losses"`
}

type GetUserFilters struct {
	Uuid          *string        `query:"uuid"`
	ChatId        *int64         `json:"chat_id"`
	TgName        *string        `json:"tg_name"`
	Balance       *types.Decimal `json:"balance"`
	WalletAddress *string        `query:"wallet_address"`
	TotalWins     *int64         `query:"total_wins"`
	TotalLosses   *int64         `query:"total_losses"`
	Limit         *int           `json:"limit"`
	Offset        *int           `json:"offset"`
}
