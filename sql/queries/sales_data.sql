-- name: CreateSalesData :one
INSERT INTO sales_data (count, min, max, last_sold_date, agent_id)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING *;