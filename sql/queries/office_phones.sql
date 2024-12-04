-- name: CreateOfficePhone :exec
INSERT INTO office_phones (office_id, phone_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(office_id, phone_id) DO NOTHING;