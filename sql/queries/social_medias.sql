-- name: CreateSocialMedia :one
INSERT INTO social_medias (type, href, agent_id)
VALUES (
    ?,
    ?,
    ?
)
RETURNING *;