-- name: CreateAgentDesignation :exec
INSERT INTO agent_designations (agent_id, designation_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(agent_id, designation_id) DO NOTHING;