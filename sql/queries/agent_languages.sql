-- name: CreateAgentLanguage :exec
INSERT INTO agent_languages (agent_id, language_id)
VALUES (
    ?,
    ?
);