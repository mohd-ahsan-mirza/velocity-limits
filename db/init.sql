CREATE SCHEMA funds;
CREATE TABLE IF NOT EXISTS load_transaction_history
(
id BIGINT NOT NULL,
customer_id BIGINT NOT NULL,
load_amount NUMERIC NOT NULL,
transaction_time TIMESTAMPTZ NOT NULL,
created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
PRIMARY KEY(id, customer_id)
);
CREATE TABLE IF NOT EXISTS unaccepted_load_transactions_history
(
id BIGINT NOT NULL,
customer_id BIGINT NOT NULL,
load_amount NUMERIC NOT NULL,
transaction_time TIMESTAMPTZ NOT NULL,
created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
PRIMARY KEY(id, customer_id)
);
SET TIME ZONE 'UTC';