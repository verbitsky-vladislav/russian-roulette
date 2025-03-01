-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chat_id BIGINT UNIQUE NOT NULL,
    tg_name VARCHAR(255) UNIQUE NOT NULL,
    balance DECIMAL(18,8) DEFAULT 0,
    wallet_address VARCHAR(255) DEFAULT '',
    total_wins INT DEFAULT 0,
    total_losses INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
