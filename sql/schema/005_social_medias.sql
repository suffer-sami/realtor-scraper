-- +goose Up
CREATE TABLE social_medias (
    id INTEGER PRIMARY KEY,
    type TEXT,
    href TEXT,
    agent_id TEXT,
    CONSTRAINT fk_agents
    FOREIGN KEY (agent_id)
    REFERENCES agents(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE social_medias;