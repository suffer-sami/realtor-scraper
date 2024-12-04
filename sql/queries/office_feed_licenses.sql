-- name: CreateOfficeFeedLicense :exec
INSERT INTO office_feed_licenses (office_id, feed_license_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(office_id, feed_license_id) DO NOTHING;