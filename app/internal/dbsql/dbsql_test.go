package db

import (
	"testing"
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func TestIsTransactionIDDuplicateForCustomer(t *testing.T) {
	dbConnectionString := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"
	p, dbOpenErr := sql.Open("postgres", dbConnectionString)
	if dbOpenErr != nil {
		log.Fatal("Failed to open a DB connection: ", dbOpenErr)
	}
	defer p.Close()
	dbs := dbsql{db: p}
	result := dbs.IsTransactionIDDuplicateForCustomer("1","1")
	if result != false {
		t.Errorf("IsTransactionIDDuplicateForCustomer was incorrect. Should have gotten false")
	}
}