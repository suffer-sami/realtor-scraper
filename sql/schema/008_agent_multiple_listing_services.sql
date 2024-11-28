-- +goose Up
CREATE TABLE agent_multiple_listing_services (
    agent_id TEXT REFERENCES agents(id) ON DELETE CASCADE,
    multiple_listing_service_id INTEGER REFERENCES multiple_listing_services(id) ON DELETE CASCADE,
    PRIMARY KEY (agent_id, multiple_listing_service_id)
);

-- +goose Down
DROP TABLE agent_multiple_listing_services;