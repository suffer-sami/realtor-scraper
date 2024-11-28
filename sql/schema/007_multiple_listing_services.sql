-- +goose Up
CREATE TABLE multiple_listing_services (
    id INTEGER PRIMARY KEY,
    abbreviation TEXT,
    inactivation_date DATETIME,
    license_number TEXT,
    member_id TEXT,
    type TEXT,
    is_primary BOOLEAN
);

-- +goose Down
DROP TABLE multiple_listing_services;