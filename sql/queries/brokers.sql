-- name: CreateBroker :one
INSERT INTO brokers (fulfillment_id, name, photo, video)
VALUES (
    ?,
    ?,
    ?,
    ?
)
ON CONFLICT(fulfillment_id) DO NOTHING
RETURNING id;

-- name: GetBrokerID :one
SELECT id FROM brokers
WHERE fulfillment_id = ? 
LIMIT 1;