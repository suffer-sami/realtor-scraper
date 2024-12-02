-- name: CreateArea :one
INSERT INTO areas (name, state_code)
VALUES (
    ?,
    ?
)
ON CONFLICT(name, state_code) DO NOTHING
RETURNING *;

-- name: GetArea :one
SELECT * FROM areas
WHERE name = ? AND state_code = ?
LIMIT 1;