-- +goose Up
CREATE TABLE specializations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE
);

-- +goose Down
DROP TABLE specializations;