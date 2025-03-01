-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE game (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    creator_uuid UUID NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('waiting', 'active', 'finished', 'cancelled')),
    bet_amount DECIMAL(10,2) NOT NULL,
    bullet_count INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE game_players (
    game_uuid UUID REFERENCES game(uuid) ON DELETE CASCADE,
    user_uuid UUID NOT NULL,
    has_shot BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (game_uuid, user_uuid)
);

CREATE TABLE game_bets (
    game_uuid UUID REFERENCES game(uuid) ON DELETE CASCADE,
    user_uuid UUID NOT NULL,
    bet_amount DECIMAL(10,2) NOT NULL,
    PRIMARY KEY (game_uuid, user_uuid)
);

CREATE TABLE game_rounds (
    game_uuid UUID REFERENCES game(uuid) ON DELETE CASCADE,
    game_id INT REFERENCES game(uuid) ON DELETE CASCADE,
    user_id BIGINT REFERENCES users(telegram_id),
    action VARCHAR(10) CHECK (action IN ('pull', 'pass')), -- действие игрока
    result VARCHAR(10) CHECK (result IN ('miss', 'shot', NULL)), -- попал или нет
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS game_bets;
DROP TABLE IF EXISTS game_players;
DROP TABLE IF EXISTS game;
-- +goose StatementEnd
