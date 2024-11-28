-- name: CreateLanguage :one
INSERT INTO languages (name)
VALUES (?)
ON CONFLICT(name) DO UPDATE SET name = EXCLUDED.name
RETURNING *;