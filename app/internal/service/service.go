package service

import (
	"app/internal"
	"database/sql"
	"encoding/json"
	"log"
	"sort"
	"strings"
	"time"

	dbsql "app/internal/dbsql"
)

type service struct {
	db internal.Db
}

//New method
func New(db *sql.DB) internal.Service {
	return &service{db: dbsql.New(db)}
}

func (s *service) LoadFunds(request string) (bool, []byte, error) {
	request = strings.ReplaceAll(request, "$", "")
	loadTransactionRecord := internal.LoadTransactionRecord{}
	json.Unmarshal([]byte(request), &loadTransactionRecord)

	if s.db.IsTransactionIDDuplicate(loadTransactionRecord.ID) {
		return true, nil, nil
	}

	records := s.db.GetAllRecordsForLatestTransactionByCustomerID("week", loadTransactionRecord.CustomerID)
	// sorting the records by latest transaction time so the last transaction date will be the first record in the array
	sort.Slice(records, func(i, j int) bool {
		return records[i].TransactionTime.After(records[j].TransactionTime)
	})
	var lastTransactionTimeStamp time.Time
	if len(records) > 0 {
		lastTransactionTimeStamp = records[0].TransactionTime

		// A maximum of $20,000 can be loaded per week
		// A maximum of 3 loads can be performed per day, regardless of amount
		// A maximum of $5,000 can be loaded per day

	}

	result, resultErr := s.db.InsertLoadTransactionRecord(&loadTransactionRecord)
	if resultErr != nil {
		log.Fatal(resultErr)
	}

	var responseStr string
	if result {
		responseStr = `{"id":"` + loadTransactionRecord.ID +
			`","customer_id":"` + loadTransactionRecord.CustomerID + `","accepted": true}`
	} else {
		responseStr = `{"id":"` + loadTransactionRecord.ID +
			`","customer_id":"` + loadTransactionRecord.CustomerID + `","accepted": false}`
	}
	response, _ := json.Marshal(responseStr)
	return false, response, nil
}
