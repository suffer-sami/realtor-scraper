-- name: CreateZip :one
INSERT INTO zips (zip_code)
VALUES (?)
ON CONFLICT(zip_code) DO NOTHING
RETURNING id;

-- name: GetZipID :one
SELECT id FROM zips 
WHERE zip_code = ? 
LIMIT 1;