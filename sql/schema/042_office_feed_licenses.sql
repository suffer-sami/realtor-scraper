-- +goose Up
CREATE TABLE office_feed_licenses (
    office_id INTEGER,
    feed_license_id INTEGER,
    CONSTRAINT fk_offices
        FOREIGN KEY (office_id)
        REFERENCES offices(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_feed_licenses
        FOREIGN KEY (feed_license_id)
        REFERENCES feed_licenses(id)
        ON DELETE CASCADE,
    PRIMARY KEY (office_id, feed_license_id)
);

-- +goose Down
DROP TABLE office_feed_licenses;