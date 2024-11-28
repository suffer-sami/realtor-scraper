-- +goose Up
CREATE TABLE languages (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE
);

-- +goose Down
DROP TABLE languages;