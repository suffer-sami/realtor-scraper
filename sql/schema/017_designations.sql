-- +goose Up
CREATE TABLE designations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE
);

-- +goose Down
DROP TABLE designations;