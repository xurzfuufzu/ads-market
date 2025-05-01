-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS companies
(
    id           UUID PRIMARY KEY,
    name         VARCHAR(255) UNIQUE NOT NULL,
    email        VARCHAR(255) UNIQUE NOT NULL,
    password     VARCHAR(255)        NOT NULL,
    phone        VARCHAR(20),
    description  TEXT,
    account_type VARCHAR(50),
    created_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS companies;
-- +goose StatementEnd
