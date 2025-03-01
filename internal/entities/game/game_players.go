package game

//
//type GamePlayers struct {
//	GameUuid string `json:"game_uuid"`
//	UserUuid string `json:"user_uuid"`
//	HasShot bool `json:"has_shot"`
//	IsAlive bool `json:"is_alive"`
//}
//
//type CreateGamePlayers struct {
//
//}
//
//type UpdateGamePlayers struct {
//
//}
//
//type GetGamePlayersFilters struct {
//
//}
//CREATE TABLE game_players (
//game_uuid UUID REFERENCES game(uuid) ON DELETE CASCADE,
//user_uuid UUID REFERENCES users(uuid) ON DELETE CASCADE,
//has_shot BOOLEAN DEFAULT FALSE,
//is_alive BOOLEAN DEFAULT TRUE,
//PRIMARY KEY (game_uuid, user_uuid)
//);
