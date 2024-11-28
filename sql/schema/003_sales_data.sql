-- +goose Up
CREATE TABLE sales_data (
    id INTEGER PRIMARY KEY,
    count INTEGER,
    min INTEGER,
    max INTEGER,
    last_sold_date DATE,
    agent_id TEXT,
    CONSTRAINT fk_agents
    FOREIGN KEY (agent_id)
    REFERENCES agents(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE sales_data;