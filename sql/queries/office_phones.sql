-- name: CreateOfficePhone :exec
INSERT INTO office_phones (office_id, phones_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(office_id, phones_id) DO NOTHING;