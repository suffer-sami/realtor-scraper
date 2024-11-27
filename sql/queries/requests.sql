-- name: CreateRequest :one
INSERT INTO requests (created_at, updated_at, offset, results_per_page)
VALUES (
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    ?,
    ?
)
RETURNING *;

-- name: GetRequests :many
SELECT offset, results_per_page FROM requests;
