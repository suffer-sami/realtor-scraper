-- name: CreateDesignation :one
INSERT INTO designations (name)
VALUES (?)
ON CONFLICT(name) DO UPDATE SET name = EXCLUDED.name
RETURNING *;