package service

import (
	"app/internal"
	"database/sql"
	"encoding/json"
	"log"
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
	json.Unmarshal([]byte(request), &loadTransactionRecord)

	if s.db.IsTransactionIDDuplicate(loadTransactionRecord.ID) {
		return true, nil, nil
	}

	s.db.GetAllRecordsForLatestTransactionByCustomerID("week", loadTransactionRecord.CustomerID)

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
