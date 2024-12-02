-- name: CreateDesignation :one
INSERT INTO designations (name)
VALUES (?)
ON CONFLICT(name) DO NOTHING
RETURNING *;

-- name: GetDesignation :one
SELECT * FROM designations 
WHERE name = ? 
LIMIT 1;