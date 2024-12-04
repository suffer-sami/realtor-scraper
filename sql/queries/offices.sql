-- name: CreateOffice :one
INSERT INTO offices (name, photo, website, email, slogan, video, fulfillment_id, address_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(fulfillment_id) DO NOTHING
RETURNING id;

-- name: GetOfficeID :one
SELECT id FROM offices
WHERE fulfillment_id = ? 
LIMIT 1;