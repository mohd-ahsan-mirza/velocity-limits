package db

import (
	"app/internal"
	"database/sql"
	"log"
	"strconv"
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
	id, idErr := strconv.Atoi(record.ID)
	if idErr != nil {
		return false, idErr
	}
	customerID, customerIDErr := strconv.Atoi(record.CustomerID)
	if customerIDErr != nil {
		return false, customerIDErr
	}
	loadAmount, loadAmountErr := strconv.ParseFloat(record.LoadAmount, 64)
	if loadAmountErr != nil {
		return false, loadAmountErr
	}

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

func (s *dbsql) IsTransactionIDDuplicate(id string) bool {
	ID, idErr := strconv.Atoi(id)
	if idErr != nil {
		log.Fatal(idErr)
	}

	var count int

	row := s.db.QueryRow("SELECT COUNT(*) FROM load_transaction_history WHERE id = $1;", ID)
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
// Getting all records for the timeInterval of the last transaction date by customer
func (s *dbsql) GetAllRecordsForLatestTransactionByCustomerID(timeInterval string, custid string) []internal.LoadTransactionRecord {
	customerID, customerIDErr := strconv.Atoi(custid)
	if customerIDErr != nil {
		log.Fatal(customerIDErr)
	}

	sqlStatement := `SELECT
		id,
		customer_id,
		load_amount,
		transaction_time
	FROM
		load_transaction_history
	WHERE
		customer_id = $1
		AND date_trunc($2, transaction_time) = (
			SELECT
				date_trunc($2, (
						SELECT
							transaction_time FROM load_transaction_history
						WHERE
							customer_id = $1
						ORDER BY
							customer_id, transaction_time DESC
						LIMIT 1))) ORDER BY transaction_time;`
	rows, err := s.db.Query(sqlStatement, customerID, timeInterval)
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
