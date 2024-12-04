// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: broker_agents.sql

package database

import (
	"context"
	"database/sql"
)

const createBrokerAgent = `-- name: CreateBrokerAgent :exec
INSERT INTO broker_agents (agent_id, broker_id)
VALUES (
    ?,
    ?
)
ON CONFLICT(agent_id) DO NOTHING
`

type CreateBrokerAgentParams struct {
	AgentID  sql.NullString
	BrokerID sql.NullInt64
}

func (q *Queries) CreateBrokerAgent(ctx context.Context, arg CreateBrokerAgentParams) error {
	_, err := q.db.ExecContext(ctx, createBrokerAgent, arg.AgentID, arg.BrokerID)
	return err
}
