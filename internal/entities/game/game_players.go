package game

// Во всем репозитории ловить ошибки на REFERENCES, потому что game_uuid и user_uuid - это ключи в бд других таблиц

type GamePlayer struct {
	Uuid     string `json:"uuid"`
	GameUuid string `json:"game_uuid"`
	UserUuid string `json:"user_uuid"`
	Name     string `json:"name"`
	HasShot  bool   `json:"has_shot"`
	IsAlive  bool   `json:"is_alive"`
}

type CreateGamePlayer struct {
	GameUuid string `json:"game_uuid"`
	UserUuid string `json:"user_uuid"`
	Name     string `json:"name"`
}

type UpdateGamePlayer struct {
	UserUuid string `json:"user_uuid"`
	Name     string `json:"name"`
	HasShot  bool   `json:"has_shot"`
	IsAlive  bool   `json:"is_alive"`
}

type GetGamePlayersFilters struct { // todo add check on filters exists, потому что все поинтеры
	Uuid     *string `json:"uuid"`
	GameUuid *string `json:"game_uuid"`
	UserUuid *string `json:"user_uuid"`
	Name     *string `json:"name"`
	HasShot  *bool   `json:"has_shot"`
	IsAlive  *bool   `json:"is_alive"`
}
