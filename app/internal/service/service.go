package service

import (
	"app/internal"
	"database/sql"
	"encoding/json"
	"log"
	"sort"
	"strings"

	dbsql "app/internal/dbsql"
)

// MAXWEEKLYLOADLIMIT Condition 1
var MAXWEEKLYLOADLIMIT float64 = 20000

// MAXDAILYLOADLIMIT Condition 2
var MAXDAILYLOADLIMIT float64 = 5000

// MAXDAILYTRANSACTIONLIMIT Condition 3
var MAXDAILYTRANSACTIONLIMIT int = 3

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
	if s.db.IsTransactionIDDuplicateForCustomer(loadTransactionRecord.ID, loadTransactionRecord.CustomerID) ||
		s.db.IsTransactionIDInUnacceptableTransacationLoadHistoryForCustomer(loadTransactionRecord.ID, loadTransactionRecord.CustomerID) {
		return true, nil, nil
	}

	unacceptedTransactionRespStr := `{"id":"` + loadTransactionRecord.ID +
		`","customer_id":"` + loadTransactionRecord.CustomerID + `","accepted": false}`
	response, _ := json.Marshal(unacceptedTransactionRespStr)

	// Getting all transactions that happened within the week of the transaction date by customer
	records := s.db.GetAllRecordsForTransactionTimeByCustomerID("week", loadTransactionRecord.CustomerID, loadTransactionRecord.TransactionTime)
	// sorting the records by latest transaction time so the last transaction date will be the first record in the array
	sort.Slice(records, func(i, j int) bool {
		return records[i].TransactionTime.After(records[j].TransactionTime)
	})

	// If the customer is loading funds the first time only need to check the daily number of transactions limit condition
	if len(records) > 0 {
		latestTransactionTimeStamp := loadTransactionRecord.TransactionTime
		var allTransactionRecordsOfLoadTransactionDate []internal.LoadTransactionRecord

		for _, record := range records {
			if internal.DateEqual(latestTransactionTimeStamp, record.TransactionTime) {
				allTransactionRecordsOfLoadTransactionDate = append(allTransactionRecordsOfLoadTransactionDate, record)
			}
		}

		// A maximum of $20,000 can be loaded per week
		totalLoadedAmountOfTheWeek := internal.SumUpLoadAmountsofTransactionRecords(allTransactionRecordsOfLoadTransactionDate)
		if (totalLoadedAmountOfTheWeek + internal.ParseFloat(loadTransactionRecord.LoadAmount)) > MAXWEEKLYLOADLIMIT {
			_, resultErr := s.db.InsertUnacceptedLoadTransactionRecord(&loadTransactionRecord)
			if resultErr != nil {
				log.Fatal(resultErr)
			}
			return false, response, nil
		}
		// A maximum of 3 loads can be performed per day, regardless of amount
		if len(allTransactionRecordsOfLoadTransactionDate) >= MAXDAILYTRANSACTIONLIMIT {
			_, resultErr := s.db.InsertUnacceptedLoadTransactionRecord(&loadTransactionRecord)
			if resultErr != nil {
				log.Fatal(resultErr)
			}
			return false, response, nil
		}
		// A maximum of $5,000 can be loaded per day
		totalLoadedAmountOftheDay := internal.SumUpLoadAmountsofTransactionRecords(allTransactionRecordsOfLoadTransactionDate)
		if (totalLoadedAmountOftheDay + internal.ParseFloat(loadTransactionRecord.LoadAmount)) > MAXDAILYLOADLIMIT {
			_, resultErr := s.db.InsertUnacceptedLoadTransactionRecord(&loadTransactionRecord)
			if resultErr != nil {
				log.Fatal(resultErr)
			}
			return false, response, nil
		}
	} else {
		if (internal.ParseFloat(loadTransactionRecord.LoadAmount)) > MAXDAILYLOADLIMIT {
			_, resultErr := s.db.InsertUnacceptedLoadTransactionRecord(&loadTransactionRecord)
			if resultErr != nil {
				log.Fatal(resultErr)
			}
			return false, response, nil
		}
	}

	//Insert the record
	result, resultErr := s.db.InsertLoadTransactionRecord(&loadTransactionRecord)
	if resultErr != nil {
		log.Fatal(resultErr)
	}
	if result {
		responseStr := `{"id":"` + loadTransactionRecord.ID +
			`","customer_id":"` + loadTransactionRecord.CustomerID + `","accepted": true}`
		response, _ = json.Marshal(responseStr)
	}
	return false, response, nil
}
