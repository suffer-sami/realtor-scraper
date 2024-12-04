-- +goose Up
ALTER TABLE agents
RENAME COLUMN web_url TO profile_url;

ALTER TABLE agents
RENAME COLUMN href TO website;

-- +goose Down
ALTER TABLE agents
RENAME COLUMN profile_url TO web_url;

ALTER TABLE agents
RENAME COLUMN website TO href;