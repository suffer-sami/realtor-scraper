-- name: CreateAgentZip :exec
INSERT INTO agent_zips (agent_id, zip_id)
VALUES (
    ?,
    ?
);