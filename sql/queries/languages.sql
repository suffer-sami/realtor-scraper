-- name: CreateLanguage :one
INSERT INTO languages (name)
VALUES (?)
ON CONFLICT(name) DO NOTHING
RETURNING *;

-- name: GetLanguage :one
SELECT * FROM languages 
WHERE name = ? 
LIMIT 1;