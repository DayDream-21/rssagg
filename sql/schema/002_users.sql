-- +goose Up

CREATE TABLE rssagg.users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down

DROP TABLE rssagg.users;