-- name: CreateAgentMarketingArea :exec
INSERT INTO agent_marketing_areas (agent_id, area_id)
VALUES (
    ?,
    ?
);