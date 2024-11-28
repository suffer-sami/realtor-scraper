-- name: CreateListingsData :one
INSERT INTO listings_data (count, min, max, last_listing_date, agent_id)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING *;