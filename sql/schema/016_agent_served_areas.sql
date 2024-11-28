-- +goose Up
CREATE TABLE agent_served_areas (
    agent_id TEXT,
    area_id INTEGER,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_areas
        FOREIGN KEY (area_id)
        REFERENCES areas(id)
        ON DELETE CASCADE,
    PRIMARY KEY (agent_id, area_id)
);

-- +goose Down
DROP TABLE agent_served_areas;