-- +goose Up
CREATE TABLE requests (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    offset INTEGER NOT NULL,
    results_per_page INTEGER NOT NULL,
    UNIQUE (offset, results_per_page)
);

-- +goose Down
DROP TABLE requests;