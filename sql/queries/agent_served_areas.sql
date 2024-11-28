-- name: CreateAgentServedArea :exec
INSERT INTO agent_served_areas (agent_id, area_id)
VALUES (
    ?,
    ?
);