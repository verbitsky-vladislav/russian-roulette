package game

import "time"

type GameAction string

var (
	Pull GameAction = "pull"
	Pass GameAction = "pass"
)

type GameActionResult string

var (
	Miss GameActionResult = "miss"
	Show GameActionResult = "show"
)

type GameRound struct {
	Uuid     string           `json:"uuid"`
	GameUuid string           `json:"game_uuid"`
	UserUuid string           `json:"user_uuid"`
	Action   GameAction       `json:"action"`
	Result   GameActionResult `json:"result"`
	Created  time.Time        `json:"created"`
}

type CreateGameRound struct {
	GameUuid string           `json:"game_uuid"`
	UserUuid string           `json:"user_uuid"`
	Action   GameAction       `json:"action"`
	Result   GameActionResult `json:"result"`
}

type GetGameRoundsFilters struct { // todo add check if filters null return nill
	Uuid     *string           `json:"uuid"`
	GameUuid *string           `json:"game_uuid"`
	UserUuid *string           `json:"user_uuid"`
	Action   *GameAction       `json:"action"`
	Result   *GameActionResult `json:"result"`
}
