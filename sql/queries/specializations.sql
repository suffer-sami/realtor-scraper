-- name: CreateSpecialization :one
INSERT INTO specializations (name)
VALUES (?)
ON CONFLICT(name) DO UPDATE SET name = EXCLUDED.name
RETURNING *;