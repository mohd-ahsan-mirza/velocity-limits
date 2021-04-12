package db

import (
	"app/internal"
	"database/sql"
	"strconv"
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
