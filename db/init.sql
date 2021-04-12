CREATE SCHEMA funds;
CREATE TABLE IF NOT EXISTS load_transaction_history
(
id BIGINT PRIMARY KEY NOT NULL,
customer_id BIGINT NOT NULL,
load_amount NUMERIC NOT NULL,
transaction_time TIMESTAMPTZ NOT NULL,
created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
SET TIME ZONE 'UTC';
/*
SELECT
	*
FROM
	load_transaction_history
WHERE
	customer_id = 834
	AND date_trunc('week', transaction_time) = (
		SELECT
			date_trunc('week', (
					SELECT
						transaction_time FROM load_transaction_history
					WHERE
						customer_id = 834
					ORDER BY
						customer_id, transaction_time DESC
					LIMIT 1)));
*/