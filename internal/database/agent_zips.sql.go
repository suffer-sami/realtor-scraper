// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: agent_zips.sql

package database

import (
	"context"
	"database/sql"
)

const createAgentZip = `-- name: CreateAgentZip :exec
INSERT INTO agent_zips (agent_id, zip_id)
VALUES (
    ?,
    ?
)
`

type CreateAgentZipParams struct {
	AgentID sql.NullString
	ZipID   sql.NullInt64
}

func (q *Queries) CreateAgentZip(ctx context.Context, arg CreateAgentZipParams) error {
	_, err := q.db.ExecContext(ctx, createAgentZip, arg.AgentID, arg.ZipID)
	return err
}