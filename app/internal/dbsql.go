package internal

import "time"

//LoadTransactionRecord for db records
type LoadTransactionRecord struct {
	ID              string    `json:"id"`
	CustomerID      string    `json:"customer_id"`
	LoadAmount      string    `json:"load_amount"`
	TransactionTime time.Time `json:"time"`
}

// Db interface
type Db interface {
	InsertLoadTransactionRecord(*LoadTransactionRecord) (bool, error)
	GetAllRecordsForTransactionTimeByCustomerID(string, string, time.Time) []LoadTransactionRecord
	IsTransactionIDDuplicateForCustomer(id string, custID string) bool
}
