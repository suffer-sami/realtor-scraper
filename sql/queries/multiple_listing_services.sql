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
RETURNING *;