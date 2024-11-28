-- +goose Up
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

-- +goose Down
DROP TABLE feed_licenses;