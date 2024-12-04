-- +goose Up
CREATE TABLE offices (
    id INTEGER PRIMARY KEY,
    name TEXT,
    photo TEXT,
    website TEXT,
    email TEXT,
    slogan TEXT,
    video TEXT,
    fulfillment_id INTEGER,
    address_id INTEGER,
    CONSTRAINT fk_addresses
        FOREIGN KEY (address_id) 
        REFERENCES addresses(id),
    CONSTRAINT unique_offices_fulfillment_id
        UNIQUE(fulfillment_id)
);

-- +goose Down
DROP TABLE offices;