-- name: CreateZip :one
INSERT INTO zips (zip_code)
VALUES (?)
ON CONFLICT(zip_code) DO UPDATE SET zip_code = EXCLUDED.zip_code
RETURNING *;