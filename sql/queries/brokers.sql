-- name: CreateBroker :one
INSERT INTO brokers (fulfillment_id, name, photo, video)
VALUES (
    ?,
    ?,
    ?,
    ?
)
ON CONFLICT(fulfillment_id) DO NOTHING
RETURNING *;

-- name: GetBroker :one
SELECT * FROM brokers
WHERE fulfillment_id = ? 
LIMIT 1;