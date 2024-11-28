-- name: CreateAgentUserLanguage :exec
INSERT INTO agent_user_languages (agent_id, language_id)
VALUES (
    ?,
    ?
);