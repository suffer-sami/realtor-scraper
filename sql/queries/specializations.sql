-- name: CreateSpecialization :one
INSERT INTO specializations (name)
VALUES (?)
ON CONFLICT(name) DO NOTHING
RETURNING *;

-- name: GetSpecialization :one
SELECT * FROM specializations 
WHERE name = ? 
LIMIT 1;