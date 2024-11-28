-- +goose Up
CREATE TABLE listings_data (
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

-- +goose Down
DROP TABLE listings_data;