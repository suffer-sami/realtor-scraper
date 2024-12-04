-- name: CreateAgentPhone :exec
INSERT INTO agent_phones (agent_id, phones_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(agent_id, phones_id) DO NOTHING;