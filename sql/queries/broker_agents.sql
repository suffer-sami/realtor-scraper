-- name: CreateBrokerAgent :exec
INSERT INTO broker_agents (agent_id, broker_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(agent_id) DO NOTHING;