-- +goose Up
CREATE TABLE agent_languages (
    agent_id TEXT,
    language_id INTEGER,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_languages
        FOREIGN KEY (language_id)
        REFERENCES languages(id)
        ON DELETE CASCADE,
    PRIMARY KEY (agent_id, language_id)
);

-- +goose Down
DROP TABLE agent_languages;