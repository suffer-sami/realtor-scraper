-- name: CreateAgentMarketingArea :exec
INSERT INTO agent_marketing_areas (agent_id, area_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(agent_id, area_id) DO NOTHING;