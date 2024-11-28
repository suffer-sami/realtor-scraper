-- name: CreateFeedLicense :one
INSERT INTO feed_licenses (country, license_number, state_code, agent_id)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING *;