-- +goose Up
CREATE TABLE zips (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    zip_code TEXT UNIQUE
);

-- +goose Down
DROP TABLE zips;