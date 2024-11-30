-- name: CreateListingsData :one
INSERT INTO listings_data (count, min, max, last_listing_date, agent_id)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?
)
ON CONFLICT (agent_id) 
DO UPDATE SET
    count = EXCLUDED.count,
    min = EXCLUDED.min,
    max = EXCLUDED.max,
    last_listing_date = EXCLUDED.last_listing_date
RETURNING *;