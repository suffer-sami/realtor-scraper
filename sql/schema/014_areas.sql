-- +goose Up
CREATE TABLE areas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    state_code TEXT,
    UNIQUE (name, state_code)
);

-- +goose Down
DROP TABLE areas;