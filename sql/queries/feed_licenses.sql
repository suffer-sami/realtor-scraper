-- name: CreateFeedLicense :exec
INSERT INTO feed_licenses (agent_id, country, state_code, license_number)
VALUES (
    ?,
    ?,
    ?,
    ?
)
ON CONFLICT(agent_id, country, state_code, license_number) DO NOTHING;