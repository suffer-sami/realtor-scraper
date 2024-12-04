-- name: CreateDesignation :one
INSERT INTO designations (name)
VALUES (?)
ON CONFLICT(name) DO NOTHING
RETURNING id;

-- name: GetDesignationID :one
SELECT id FROM designations 
WHERE name = ? 
LIMIT 1;