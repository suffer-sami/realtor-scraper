-- +goose Up
CREATE TEMPORARY TABLE temp_social_medias AS SELECT * FROM social_medias;

DROP TABLE social_medias;

CREATE TABLE social_medias (
    id INTEGER PRIMARY KEY,
    type TEXT,
    href TEXT,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE,
    CONSTRAINT unique_agent_href 
        UNIQUE (agent_id, href)
);

INSERT INTO social_medias SELECT * FROM temp_social_medias;

DROP TABLE temp_social_medias;

-- +goose Down
CREATE TEMPORARY TABLE temp_social_medias AS SELECT * FROM social_medias;

DROP TABLE social_medias;

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

INSERT INTO social_medias SELECT * FROM temp_social_medias;

DROP TABLE temp_social_medias;