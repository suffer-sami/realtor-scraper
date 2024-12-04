-- +goose Up
ALTER TABLE agent_phones
RENAME COLUMN phones_id TO phone_id;

ALTER TABLE office_phones
RENAME COLUMN phones_id TO phone_id;

-- +goose Down
ALTER TABLE agent_phones
RENAME COLUMN phone_id TO phones_id;

ALTER TABLE office_phones
RENAME COLUMN phone_id TO phones_id;