-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ads
(
    id            UUID PRIMARY KEY,
    title         TEXT      NOT NULL,
    company_name  TEXT      NOT NULL,
    description   TEXT,
    priceFrom     INTEGER,
    priceTo       INTEGER,
    status        TEXT      NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    platforms     TEXT[]    NOT NULL,
    category      TEXT,
    target_city   TEXT,
    responses_count INTEGER DEFAULT 0,

    FOREIGN KEY (company_name) REFERENCES companies (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ads;
-- +goose StatementEnd