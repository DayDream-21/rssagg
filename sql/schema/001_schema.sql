-- +goose Up

CREATE SCHEMA IF NOT EXISTS rssagg;

-- +goose Down

DROP SCHEMA IF EXISTS rssagg;