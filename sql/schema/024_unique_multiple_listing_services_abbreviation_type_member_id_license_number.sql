-- +goose Up
CREATE TEMPORARY TABLE temp_multiple_listing_services AS SELECT * FROM multiple_listing_services;

DROP TABLE multiple_listing_services;

CREATE TABLE multiple_listing_services (
    id INTEGER PRIMARY KEY,
    abbreviation TEXT,
    inactivation_date DATETIME,
    license_number TEXT,
    member_id TEXT,
    type TEXT,
    is_primary BOOLEAN,
    CONSTRAINT unique_multiple_listing_services_abbreviation_type_member_id_license_number
        UNIQUE (abbreviation, type, member_id, license_number)
);

INSERT INTO multiple_listing_services SELECT * FROM temp_multiple_listing_services;

DROP TABLE temp_multiple_listing_services;

-- +goose Down
CREATE TEMPORARY TABLE temp_multiple_listing_services AS SELECT * FROM multiple_listing_services;

DROP TABLE multiple_listing_services;

CREATE TABLE multiple_listing_services (
    id INTEGER PRIMARY KEY,
    abbreviation TEXT,
    inactivation_date DATETIME,
    license_number TEXT,
    member_id TEXT,
    type TEXT,
    is_primary BOOLEAN
);

INSERT INTO multiple_listing_services SELECT * FROM temp_multiple_listing_services;

DROP TABLE temp_multiple_listing_services;