-- +goose Up
-- +goose StatementBegin
CREATE TABLE game (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    creator_uuid UUID REFERENCES users(uuid) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL CHECK (status IN ('waiting', 'active', 'finished', 'cancelled')),
    bet_amount DECIMAL(10,2) NOT NULL,
    bullet_count INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE game_players (
    game_uuid UUID REFERENCES game(uuid) ON DELETE CASCADE,
    user_uuid UUID REFERENCES users(uuid) ON DELETE CASCADE,
    has_shot BOOLEAN DEFAULT FALSE,
    is_alive BOOLEAN DEFAULT TRUE,
    PRIMARY KEY (game_uuid, user_uuid)
);

CREATE TABLE game_rounds (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    game_uuid UUID REFERENCES game(uuid) ON DELETE CASCADE,
    user_uuid UUID REFERENCES users(uuid),
    action VARCHAR(10) CHECK (action IN ('pull', 'pass')),
    result VARCHAR(10) CHECK (result IN ('miss', 'shot', NULL)),
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS game_bets;
DROP TABLE IF EXISTS game_players;
DROP TABLE IF EXISTS game;
-- +goose StatementEnd
