-- name: CreateSalesData :exec
INSERT INTO sales_data (count, min, max, last_sold_date, agent_id)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?
)
ON CONFLICT (agent_id) 
DO UPDATE SET
    count = EXCLUDED.count,
    min = EXCLUDED.min,
    max = EXCLUDED.max,
    last_sold_date = EXCLUDED.last_sold_date;