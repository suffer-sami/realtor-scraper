-- name: CreateAgentDesignation :exec
INSERT INTO agent_designations (agent_id, designation_id)
VALUES (
    ?,
    ?
);