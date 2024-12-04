-- name: CreateRawAgent :exec
INSERT INTO raw_agents (data, agent_id)
VALUES (
    ?,
    ?
)
ON CONFLICT (agent_id) 
DO UPDATE SET data = EXCLUDED.data;