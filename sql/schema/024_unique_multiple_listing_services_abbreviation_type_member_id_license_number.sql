-- +goose Up
CREATE TABLE new_multiple_listing_services (
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

INSERT INTO new_multiple_listing_services (id, abbreviation, inactivation_date, license_number, member_id, member_id, type, is_primary)
SELECT id, abbreviation, inactivation_date, license_number, member_id, member_id, type, is_primary FROM multiple_listing_services;

DROP TABLE multiple_listing_services;

ALTER TABLE new_multiple_listing_services RENAME TO multiple_listing_services;

-- +goose Down
CREATE TABLE new_multiple_listing_services (
    id INTEGER PRIMARY KEY,
    abbreviation TEXT,
    inactivation_date DATETIME,
    license_number TEXT,
    member_id TEXT,
    type TEXT,
    is_primary BOOLEAN
);

INSERT INTO new_multiple_listing_services (id, abbreviation, inactivation_date, license_number, member_id, member_id, type, is_primary)
SELECT id, abbreviation, inactivation_date, license_number, member_id, member_id, type, is_primary FROM multiple_listing_services;

DROP TABLE multiple_listing_services;

ALTER TABLE new_multiple_listing_services RENAME TO multiple_listing_services;