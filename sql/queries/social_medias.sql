-- name: CreateSocialMedia :one
INSERT INTO social_medias (type, href, agent_id)
VALUES (
    ?,
    ?,
    ?
)
ON CONFLICT (agent_id, href) DO NOTHING
RETURNING *;