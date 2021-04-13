package service

import (
	"app/internal"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

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
	// Load the request into the struct
	json.Unmarshal([]byte(request), &loadTransactionRecord)

	// Checking if id already exists. If true, abandon the process
	if s.db.IsTransactionIDDuplicate(loadTransactionRecord.ID) {
		return true, nil, nil
	}

	unacceptedTransactionRespStr := `{"id":"` + loadTransactionRecord.ID +
		`","customer_id":"` + loadTransactionRecord.CustomerID + `","accepted": false}`

	// Getting all transactions that happened within the week of the last transaction
	records := s.db.GetAllRecordsForLatestTransactionByCustomerID("week", loadTransactionRecord.CustomerID)
	// sorting the records by latest transaction time so the last transaction date will be the first record in the array
	sort.Slice(records, func(i, j int) bool {
		return records[i].TransactionTime.After(records[j].TransactionTime)
	})

	// If the customer is loading funds the first time none of the following checks are required
	if len(records) > 0 {
		latestTransactionTimeStamp := loadTransactionRecord.TransactionTime
		var allTransactionRecordsOfLastTransactionDate []internal.LoadTransactionRecord

		for _, record := range records {
			if DateEqual(latestTransactionTimeStamp, record.TransactionTime) {
				allTransactionRecordsOfLastTransactionDate = append(allTransactionRecordsOfLastTransactionDate, record)
				fmt.Println(allTransactionRecordsOfLastTransactionDate)
			}
		}

		// A maximum of $20,000 can be loaded per week
		// A maximum of 3 loads can be performed per day, regardless of amount
		// A maximum of $5,000 can be loaded per day
	}

	//Insert the record
	result, resultErr := s.db.InsertLoadTransactionRecord(&loadTransactionRecord)
	if resultErr != nil {
		log.Fatal(resultErr)
	}

	var responseStr string
	if result {
		responseStr = `{"id":"` + loadTransactionRecord.ID +
			`","customer_id":"` + loadTransactionRecord.CustomerID + `","accepted": true}`
	} else {
		responseStr = unacceptedTransactionRespStr
	}
	response, _ := json.Marshal(responseStr)
	return false, response, nil
}
