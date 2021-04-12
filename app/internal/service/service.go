package service

import (
	"app/internal"
	"database/sql"
	"encoding/json"
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
	result, resultErr := s.db.InsertLoadTransactionRecord(&loadTransactionRecord)

	if resultErr != nil {
		if strings.Contains(resultErr.Error(), "duplicate key value") {
			return true, nil, nil
		} else {
			return false, nil, resultErr
		}
	}

	var responseStr string
	if result {
		responseStr = `{"id":"` + loadTransactionRecord.ID +
			`","customer_id":"` + loadTransactionRecord.CustomerID + `","accepted": true}`
	} else {
		responseStr = `{"id":"` + loadTransactionRecord.ID +
			`","customer_id":"` + loadTransactionRecord.CustomerID + `","accepted": false}`
	}

	// if loadTransactionRecord.CustomerID == "800" {
	// 	customerID, _ := strconv.Atoi(loadTransactionRecord.CustomerID)
	// 	fmt.Println(s.db.GetAllRecordsForLatestTransactionByCustomerID("day", customerID))
	// }

	response, _ := json.Marshal(responseStr)
	return false, response, nil
}
