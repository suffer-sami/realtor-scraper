-- +goose Up
CREATE TABLE agent_zips (
    agent_id TEXT,
    zip_id INTEGER,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_zips
        FOREIGN KEY (zip_id)
        REFERENCES zips(id)
        ON DELETE CASCADE,
    PRIMARY KEY (agent_id, zip_id)
);

-- +goose Down
DROP TABLE agent_zips;