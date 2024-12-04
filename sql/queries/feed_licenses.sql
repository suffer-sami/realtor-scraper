-- name: CreateFeedLicense :one
INSERT INTO feed_licenses (country, state_code, license_number)
VALUES (
    ?,
    ?,
    ?
)
ON CONFLICT(country, state_code, license_number) 
DO UPDATE SET
    license_number = EXCLUDED.license_number
RETURNING id;