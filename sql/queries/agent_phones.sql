-- name: CreateAgentPhone :exec
INSERT INTO agent_phones (agent_id, phone_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(agent_id, phone_id) DO NOTHING;