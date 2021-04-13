package db

import (
	"app/internal"
	"database/sql"
	"log"
	"time"
)

type dbsql struct {
	db *sql.DB
}

//New method
func New(db *sql.DB) internal.Db {
	return &dbsql{db: db}
}

func (s *dbsql) InsertLoadTransactionRecord(record *internal.LoadTransactionRecord) (bool, error) {
	id := internal.ParseInt(record.ID)
	customerID := internal.ParseInt(record.CustomerID)
	loadAmount := internal.ParseFloat(record.LoadAmount)

	sqlStatement := `
	INSERT INTO load_transaction_history (id, customer_id, load_amount, transaction_time)
	VALUES ($1, $2, $3, $4)`
	result, err := s.db.Exec(sqlStatement, id,
		customerID, loadAmount,
		record.TransactionTime)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		return true, nil
	}
	return false, nil
}

func (s *dbsql) IsTransactionIDDuplicateForCustomer(id string, custID string) bool {
	ID := internal.ParseInt(id)
	customerID := internal.ParseInt(custID)

	var count int

	row := s.db.QueryRow("SELECT COUNT(*) FROM load_transaction_history WHERE id = $1 AND customer_id = $2;", ID, customerID)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return true
	}

	return false
}

// https://www.postgresqltutorial.com/postgresql-date_trunc/
// Getting all records for the timeInterval of the transaction date by customer
func (s *dbsql) GetAllRecordsForTransactionTimeByCustomerID(timeInterval string, custid string, latestTransactionTimeStamp time.Time) []internal.LoadTransactionRecord {
	customerID := internal.ParseInt(custid)

	sqlStatement := `SELECT
		id,
		customer_id,
		load_amount,
		transaction_time
	FROM
		load_transaction_history
	WHERE
		customer_id = $1
		AND date_trunc($2, transaction_time) = date_trunc($2, $3::timestamptz) ORDER BY transaction_time;`
	rows, err := s.db.Query(sqlStatement, customerID, timeInterval, latestTransactionTimeStamp)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var loadTransactionRecords []internal.LoadTransactionRecord
	var id string
	var custID string
	var loadAmount string
	var transactionTime string
	for rows.Next() {
		err := rows.Scan(&id, &custID, &loadAmount, &transactionTime)
		if err != nil {
			log.Fatal(err)
		}
		transactionTime, _ := time.Parse(time.RFC3339, transactionTime)

		loadTransactionRecords = append(loadTransactionRecords,
			internal.LoadTransactionRecord{ID: id, CustomerID: custID,
				LoadAmount: loadAmount, TransactionTime: transactionTime})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return loadTransactionRecords
}
