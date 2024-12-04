-- +goose Up
CREATE TABLE new_feed_licenses (
    id INTEGER PRIMARY KEY,
    country TEXT,
    license_number TEXT,
    state_code TEXT,
    CONSTRAINT unique_feed_license
        UNIQUE (country, state_code, license_number)
);

INSERT INTO new_feed_licenses (
    id,
    country,
    license_number,
    state_code
) SELECT 
    id,
    country,
    license_number,
    state_code
FROM feed_licenses;

DROP TABLE feed_licenses;

ALTER TABLE new_feed_licenses RENAME TO feed_licenses;

-- +goose Down
CREATE TABLE new_feed_licenses (
    id INTEGER PRIMARY KEY,
    country TEXT,
    license_number TEXT,
    state_code TEXT,
    agent_id TEXT,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE
    CONSTRAINT unique_feed_license
        UNIQUE (agent_id, country, state_code, license_number)
);

INSERT INTO new_feed_licenses (
    id,
    country,
    license_number,
    state_code
) SELECT 
    id,
    country,
    license_number,
    state_code
FROM feed_licenses;

DROP TABLE feed_licenses;

ALTER TABLE new_feed_licenses RENAME TO feed_licenses;