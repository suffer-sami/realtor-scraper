-- +goose Up
ALTER TABLE agents
ADD COLUMN photo TEXT;

-- +goose Down
ALTER TABLE agents
DROP COLUMN photo;