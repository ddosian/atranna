-- +goose Up
CREATE TABLE devices (
    id         TEXT PRIMARY KEY,
    name       TEXT,
    desc       TEXT,
    first_seen TEXT NOT NULL,
    last_seen  TEXT NOT NULL
);

-- +goose Down
DROP TABLE devices;
