-- +goose Up
CREATE TABLE agent_phones (
    agent_id TEXT,
    phones_id INTEGER,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_phones
        FOREIGN KEY (phones_id)
        REFERENCES phones(id)
        ON DELETE CASCADE,
    PRIMARY KEY (agent_id, phones_id)
);

-- +goose Down
DROP TABLE agent_phones;