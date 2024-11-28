-- +goose Up
CREATE TABLE agent_specializations (
    agent_id TEXT,
    specialization_id INTEGER,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_specializations
        FOREIGN KEY (specialization_id)
        REFERENCES specializations(id)
        ON DELETE CASCADE,
    PRIMARY KEY (agent_id, specialization_id)
);

-- +goose Down
DROP TABLE agent_specializations;