-- name: CreateAddress :one
INSERT INTO addresses (line, line2, city, country, postal_code, state, state_code)
VALUES (?, ?, ?, ?, ?, ?, ?)
ON CONFLICT (line, line2, city, state_code, postal_code)
    DO NOTHING
RETURNING id;

-- name: GetAddress :one
SELECT id
FROM addresses
WHERE line = ?
  AND line2 = ?
  AND city = ?
  AND state_code = ?
  AND postal_code = ?
LIMIT 1;
