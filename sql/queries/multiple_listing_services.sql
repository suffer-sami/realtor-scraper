-- name: CreateMultipleListingService :one
INSERT INTO multiple_listing_services (abbreviation, inactivation_date, license_number, member_id, type, is_primary)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
ON CONFLICT (abbreviation, license_number, member_id, type)
    DO NOTHING
RETURNING *;

-- name: GetMultipleListingService :one
SELECT * FROM multiple_listing_services
WHERE abbreviation = ? AND type = ? AND member_id = ? AND license_number = ?
LIMIT 1;

-- name: UpdateMultipleListingServiceInactivationDate :exec
UPDATE multiple_listing_services
SET inactivation_date = ?
WHERE id = ?;