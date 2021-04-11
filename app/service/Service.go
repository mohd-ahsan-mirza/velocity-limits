package service

import "database/sql"

//Service interface
type Service interface {
}

type service struct {
	db *sql.DB
}

//New method
func New(db *sql.DB) Service {
	return &service{db: db}
}
