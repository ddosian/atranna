-- +goose Up
CREATE TABLE interfaces (
    id          TEXT PRIMARY KEY,
    device_id   TEXT NOT NULL REFERENCES devices(id),  
    name        TEXT,
    speed       TEXT,
    ip          TEXT,
    mac         TEXT,
    first_seen  TEXT NOT NULL,
    last_seen   TEXT NOT NULL
);

-- +goose Down
DROP TABLE interfaces;
