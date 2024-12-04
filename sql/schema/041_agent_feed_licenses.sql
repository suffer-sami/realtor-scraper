-- +goose Up
CREATE TABLE agent_feed_licenses (
    agent_id TEXT,
    feed_license_id INTEGER,
    CONSTRAINT fk_agents
        FOREIGN KEY (agent_id)
        REFERENCES agents(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_feed_licenses
        FOREIGN KEY (feed_license_id)
        REFERENCES feed_licenses(id)
        ON DELETE CASCADE,
    PRIMARY KEY (agent_id, feed_license_id)
);

-- +goose Down
DROP TABLE agent_feed_licenses;