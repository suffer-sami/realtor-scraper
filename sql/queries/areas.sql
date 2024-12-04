-- name: CreateArea :one
INSERT INTO areas (name, state_code)
VALUES (
    ?,
    ?
)
ON CONFLICT(name, state_code) DO NOTHING
RETURNING id;

-- name: GetAreaID :one
SELECT id FROM areas
WHERE name = ? AND state_code = ?
LIMIT 1;