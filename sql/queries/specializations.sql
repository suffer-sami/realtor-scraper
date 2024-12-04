-- name: CreateSpecialization :one
INSERT INTO specializations (name)
VALUES (?)
ON CONFLICT(name) DO NOTHING
RETURNING id;

-- name: GetSpecializationID :one
SELECT id FROM specializations 
WHERE name = ? 
LIMIT 1;