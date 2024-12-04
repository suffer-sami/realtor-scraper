-- +goose Up
CREATE TABLE addresses (
    id INTEGER PRIMARY KEY,
    line TEXT,
    line2 TEXT,
    city TEXT,
    country TEXT,
    postal_code TEXT,
    state TEXT,
    state_code TEXT,
    CONSTRAINT unique_address 
        UNIQUE (line, line2, city, state_code, postal_code)
);

-- +goose Down
DROP TABLE addresses;
