-- +goose Up
CREATE TABLE agent_designations (
    agent_id TEXT,
    designation_id INTEGER,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_designations
        FOREIGN KEY (designation_id)
        REFERENCES designations(id)
        ON DELETE CASCADE,
    PRIMARY KEY (agent_id, designation_id)
);

-- +goose Down
DROP TABLE agent_designations;