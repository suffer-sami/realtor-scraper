-- +goose Up
CREATE TABLE new_sales_data (
    id INTEGER PRIMARY KEY,
    count INTEGER,
    min INTEGER,
    max INTEGER,
    last_sold_date DATETIME,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE
    CONSTRAINT unique_sales_data_agent_id
        UNIQUE (agent_id)
);

INSERT INTO new_sales_data (id, count, min, max, last_sold_date, agent_id)
SELECT id, count, min, max, last_sold_date, agent_id FROM sales_data;

DROP TABLE sales_data;

ALTER TABLE new_sales_data RENAME TO sales_data;

-- +goose Down
CREATE TABLE new_sales_data (
    id INTEGER PRIMARY KEY,
    count INTEGER,
    min INTEGER,
    max INTEGER,
    last_sold_date DATETIME,
    agent_id TEXT,
    CONSTRAINT fk_agents 
        FOREIGN KEY (agent_id) 
        REFERENCES agents(id) 
        ON DELETE CASCADE
);

INSERT INTO new_sales_data (id, count, min, max, last_sold_date, agent_id)
SELECT id, count, min, max, last_sold_date, agent_id FROM sales_data;

DROP TABLE sales_data;

ALTER TABLE new_sales_data RENAME TO sales_data;