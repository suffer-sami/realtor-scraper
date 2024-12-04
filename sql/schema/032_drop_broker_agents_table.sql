-- +goose Up
DROP TABLE broker_agents;

-- +goose Down
CREATE TABLE broker_agents (
    broker_id INTEGER,
    agent_id TEXT,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_brokers
        FOREIGN KEY (broker_id)
        REFERENCES brokers(id)
        ON DELETE CASCADE,
    CONSTRAINT unique_broker_agents_agent_id
        UNIQUE (agent_id),
    PRIMARY KEY (agent_id, broker_id)
);