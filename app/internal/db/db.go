package db

import (
	"app/internal"
	"database/sql"
)

type dbsql struct {
	db *sql.DB
}

//New method
func New(db *sql.DB) internal.Db {
	return &dbsql{db: db}
}

func (s *dbsql) InsertLoadTransactionRecord(*internal.LoadTransactionRecord) (bool, error) {
	return true, nil
}
