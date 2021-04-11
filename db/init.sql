CREATE SCHEMA funds;
CREATE TABLE IF NOT EXISTS load_transaction_history
(
id BIGINT PRIMARY KEY NOT NULL,
customer_id BIGINT NOT NULL,
load_amount NUMERIC(5,2) NOT NULL,
transaction_time TIMESTAMPTZ NOT NULL,
created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
SET TIME ZONE 'UTC';
-- select date_trunc('week', current_date) as current_week