package internal

import "time"

//LoadTransactionRecord for db records
type LoadTransactionRecord struct {
	id              int
	customerID      int
	loadAmount      float64
	transactionTime time.Time
}

// Db interface
type Db interface {
	InsertLoadTransactionRecord(*LoadTransactionRecord) (bool, error)
}
