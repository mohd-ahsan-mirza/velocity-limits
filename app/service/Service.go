package service

import (
	"database/sql"
	"fmt"
)

//Service interface
type Service interface {
}

type service struct {
	db *sql.DB
}

//New method
func New(db *sql.DB) Service {
	fmt.Println("COMING IN NEW METHOD")
	return &service{db: db}
}
