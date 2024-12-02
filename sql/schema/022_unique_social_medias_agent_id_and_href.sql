-- +goose Up
CREATE TABLE new_social_medias (
    id INTEGER PRIMARY KEY,
    type TEXT,
    href TEXT,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE,
    CONSTRAINT unique_social_medias_agent_id_and_href 
        UNIQUE (agent_id, href)
);

INSERT INTO new_social_medias (id, type, href, agent_id)
SELECT id, type, href, agent_id FROM social_medias;

DROP TABLE social_medias;

ALTER TABLE new_social_medias RENAME TO social_medias;

-- +goose Down
CREATE TABLE new_social_medias (
    id INTEGER PRIMARY KEY,
    type TEXT,
    href TEXT,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE
);

INSERT INTO new_social_medias (id, type, href, agent_id)
SELECT id, type, href, agent_id FROM social_medias;

DROP TABLE social_medias;

ALTER TABLE new_social_medias RENAME TO social_medias;