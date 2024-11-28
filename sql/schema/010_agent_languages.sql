-- +goose Up
CREATE TABLE agent_languages (
    agent_id TEXT REFERENCES agents(id) ON DELETE CASCADE,
    language_id INTEGER REFERENCES languages(id),
    PRIMARY KEY (agent_id, language_id)
);

-- +goose Down
DROP TABLE agent_languages;