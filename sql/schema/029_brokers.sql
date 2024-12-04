-- +goose Up
CREATE TABLE brokers (
    id INTEGER PRIMARY KEY,
    fulfillment_id INTEGER,
    name TEXT,
    photo TEXT,
    video TEXT,
    CONSTRAINT unique_brokers_fulfillment_id
        UNIQUE(fulfillment_id)
);

-- +goose Down
DROP TABLE brokers;