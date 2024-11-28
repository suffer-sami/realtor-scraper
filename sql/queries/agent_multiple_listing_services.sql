-- name: CreateAgentMultipleListingService :exec
INSERT INTO agent_multiple_listing_services (agent_id, multiple_listing_service_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(agent_id, multiple_listing_service_id) DO NOTHING;