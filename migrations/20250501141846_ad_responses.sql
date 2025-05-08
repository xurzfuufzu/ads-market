-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ad_responses
(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ad_id         UUID NOT NULL,
    influencer_id UUID NOT NULL,
    message       TEXT,
    status        TEXT        DEFAULT 'pending', -- pending / accepted / rejected
    created_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (ad_id) REFERENCES ads (id) ON DELETE CASCADE,
    FOREIGN KEY (influencer_id) REFERENCES influencers (id) ON DELETE CASCADE,
    UNIQUE (ad_id, influencer_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ad_responses;
-- +goose StatementEnd
