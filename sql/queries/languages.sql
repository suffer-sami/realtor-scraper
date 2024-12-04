-- name: CreateLanguage :one
INSERT INTO languages (name)
VALUES (?)
ON CONFLICT(name) DO NOTHING
RETURNING id;

-- name: GetLanguageID :one
SELECT id FROM languages 
WHERE name = ? 
LIMIT 1;