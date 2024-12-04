-- name: CreatePhone :one
INSERT INTO phones (ext, number, type, is_valid)
VALUES (
    ?,
    ?,
    ?,
    ?
)
ON CONFLICT(ext, number, type) DO NOTHING
RETURNING id;

-- name: GetPhoneID :one
SELECT id FROM phones
WHERE ext = ? AND number = ? AND type = ? 
LIMIT 1;