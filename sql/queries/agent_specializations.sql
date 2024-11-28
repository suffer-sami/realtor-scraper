-- name: CreateAgentSpecialization :exec
INSERT INTO agent_specializations (agent_id, specialization_id)
VALUES (
    ?,
    ?
);