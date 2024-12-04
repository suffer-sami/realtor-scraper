// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: office_phones.sql

package database

import (
	"context"
	"database/sql"
)

const createOfficePhone = `-- name: CreateOfficePhone :exec
INSERT INTO office_phones (office_id, phones_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(office_id, phones_id) DO NOTHING
`

type CreateOfficePhoneParams struct {
	OfficeID sql.NullInt64
	PhonesID sql.NullInt64
}

func (q *Queries) CreateOfficePhone(ctx context.Context, arg CreateOfficePhoneParams) error {
	_, err := q.db.ExecContext(ctx, createOfficePhone, arg.OfficeID, arg.PhonesID)
	return err
}