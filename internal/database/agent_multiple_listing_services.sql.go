// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: agent_multiple_listing_services.sql

package database

import (
	"context"
	"database/sql"
)

const createAgentMultipleListingService = `-- name: CreateAgentMultipleListingService :exec
INSERT INTO agent_multiple_listing_services (agent_id, multiple_listing_service_id)
VALUES (
    ?,
    ?
)
`

type CreateAgentMultipleListingServiceParams struct {
	AgentID                  sql.NullString
	MultipleListingServiceID sql.NullInt64
}

func (q *Queries) CreateAgentMultipleListingService(ctx context.Context, arg CreateAgentMultipleListingServiceParams) error {
	_, err := q.db.ExecContext(ctx, createAgentMultipleListingService, arg.AgentID, arg.MultipleListingServiceID)
	return err
}