-- +goose Up
CREATE TABLE new_listings_data (
    id INTEGER PRIMARY KEY,
    count INTEGER,
    min INTEGER,
    max INTEGER,
    last_listing_date DATETIME,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE
    CONSTRAINT unique_listings_data_agent_id
        UNIQUE (agent_id)
);

INSERT INTO new_listings_data (id, count, min, max, last_listing_date, agent_id)
SELECT id, count, min, max, last_listing_date, agent_id FROM listings_data;

DROP TABLE listings_data;

ALTER TABLE new_listings_data RENAME TO listings_data;

-- +goose Down
CREATE TABLE new_listings_data (
    id INTEGER PRIMARY KEY,
    count INTEGER,
    min INTEGER,
    max INTEGER,
    last_listing_date DATETIME,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE
);

INSERT INTO new_listings_data (id, count, min, max, last_listing_date, agent_id)
SELECT id, count, min, max, last_listing_date, agent_id FROM listings_data;

DROP TABLE listings_data;

ALTER TABLE new_listings_data RENAME TO listings_data;