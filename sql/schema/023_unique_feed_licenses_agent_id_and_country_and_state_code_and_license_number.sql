-- +goose Up
CREATE TEMPORARY TABLE temp_feed_licenses AS SELECT * FROM feed_licenses;

DROP TABLE feed_licenses;

CREATE TABLE feed_licenses (
    id INTEGER PRIMARY KEY,
    country TEXT,
    license_number TEXT,
    state_code TEXT,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE,
    CONSTRAINT unique_feed_licenses_agent_id_and_country_and_state_code_and_license_number
        UNIQUE (agent_id, country, state_code, license_number)
);

INSERT INTO feed_licenses SELECT * FROM temp_feed_licenses;

DROP TABLE temp_feed_licenses;

-- +goose Down
CREATE TEMPORARY TABLE temp_feed_licenses AS SELECT * FROM feed_licenses;

DROP TABLE feed_licenses;

CREATE TABLE feed_licenses (
    id INTEGER PRIMARY KEY,
    country TEXT,
    license_number TEXT,
    state_code TEXT,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE
);

INSERT INTO feed_licenses SELECT * FROM temp_feed_licenses;

DROP TABLE temp_feed_licenses;