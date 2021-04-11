package service

import (
	"app/internal"
	"database/sql"
	"encoding/json"
	"log"
)

type service struct {
	db *sql.DB
}

//New method
func New(db *sql.DB) internal.Service {
	return &service{db: db}
}

func (s *service) LoadFunds(request string) ([]byte, error) {
	response, responseErr := json.Marshal(request)
	if responseErr != nil {
		log.Fatal(responseErr)
	}
	return response, nil
}
