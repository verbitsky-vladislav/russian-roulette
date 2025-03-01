package game

// Во всем репозитории ловить ошибки на REFERENCES, потому что game_uuid и user_uuid - это ключи в бд других таблиц

type GamePlayers struct {
	Uuid     string `json:"uuid"`
	GameUuid string `json:"game_uuid"`
	UserUuid string `json:"user_uuid"`
	HasShot  bool   `json:"has_shot"`
	IsAlive  bool   `json:"is_alive"`
}

type CreateGamePlayers struct {
	GameUuid string `json:"game_uuid"`
	UserUuid string `json:"user_uuid"`
}

type UpdateGamePlayers struct {
	Uuid    string `json:"uuid"`
	HasShot bool   `json:"has_shot"`
	IsAlive bool   `json:"is_alive"`
}

type GetGamePlayersFilters struct { // todo add check on filters exists, потому что все поинтеры
	Uuid     *string `json:"uuid"`
	GameUuid *string `json:"game_uuid"`
	UserUuid *string `json:"user_uuid"`
	HasShot  *bool   `json:"has_shot"`
	IsAlive  *bool   `json:"is_alive"`
}
