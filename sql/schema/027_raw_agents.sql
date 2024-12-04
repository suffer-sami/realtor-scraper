-- +goose Up
CREATE TABLE raw_agents (
    id INTEGER PRIMARY KEY,
    data TEXT,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE,
    CONSTRAINT unique_raw_agents_agent_id
        UNIQUE (agent_id)
);

-- +goose Down
DROP TABLE raw_agents;