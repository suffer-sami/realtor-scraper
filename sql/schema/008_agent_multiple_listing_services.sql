-- +goose Up
CREATE TABLE agent_multiple_listing_services (
    agent_id TEXT,
    multiple_listing_service_id INTEGER,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_multiple_listing_services
        FOREIGN KEY (multiple_listing_service_id)
        REFERENCES multiple_listing_services(id)
        ON DELETE CASCADE,
    PRIMARY KEY (agent_id, multiple_listing_service_id)
);

-- +goose Down
DROP TABLE agent_multiple_listing_services;