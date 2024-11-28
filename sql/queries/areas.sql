-- name: CreateArea :one
INSERT INTO areas (name, state_code)
VALUES (
    ?,
    ?
)
ON CONFLICT(name, state_code) DO UPDATE SET name = EXCLUDED.name
RETURNING *;