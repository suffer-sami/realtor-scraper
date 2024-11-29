-- +goose Up

ALTER TABLE sales_data ADD COLUMN temp_last_sold_date DATETIME;

UPDATE sales_data SET temp_last_sold_date = last_sold_date;

ALTER TABLE sales_data DROP COLUMN last_sold_date;

ALTER TABLE sales_data RENAME COLUMN temp_last_sold_date TO last_sold_date;

-- +goose Down

ALTER TABLE sales_data ADD COLUMN temp_last_sold_date DATE;

UPDATE sales_data SET temp_last_sold_date = last_sold_date;

ALTER TABLE sales_data DROP COLUMN last_sold_date;

ALTER TABLE sales_data RENAME COLUMN temp_last_sold_date TO last_sold_date;
