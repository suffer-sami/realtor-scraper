-- name: CreateAgentFeedLicense :exec
INSERT INTO agent_feed_licenses (agent_id, feed_license_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(agent_id, feed_license_id) DO NOTHING;