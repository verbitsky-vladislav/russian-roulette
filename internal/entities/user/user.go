package user

import (
	"time"
)

type User struct {
	Uuid      string    `json:"uuid"`
	ChatId    int64     `json:"chat_id"`
	TgName    string    `json:"tg_name"`
	IsPremium bool      `json:"is_premium"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUser struct {
	ChatId int64  `json:"chat_id"`
	TgName string `json:"tg_name"`
}

type UpdateUser struct {
	ChatId int64   `json:"chat_id"`
	TgName *string `json:"tg_name"`
}

type GetUserFilters struct {
	ChatId *int64  `json:"chat_id"`
	TgName *string `json:"tg_name"`
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
}
