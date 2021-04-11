package main

import (
	"bufio"
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	dbConnectionString := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	db, dbOpenErr := sql.Open("postgres", dbConnectionString)
	if dbOpenErr != nil {
		log.Fatal("Failed to open a DB connection: ", dbOpenErr)
	}
	defer db.Close()

	ctx := context.Background()
	tx, txErr := db.BeginTx(ctx, nil)
	if txErr != nil {
		log.Fatal(txErr)
	}

	rows, qCtxErr := tx.QueryContext(ctx, "SELECT NOW();")
	if qCtxErr != nil {
		log.Fatal(qCtxErr)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		println(name)
	}

	file, fileOpenErr := os.Open(os.Getenv("INPUT_FILE"))
	if fileOpenErr != nil {
		log.Fatal(fileOpenErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())

	}
	if scannerError := scanner.Err(); scannerError != nil {
		log.Fatal(scannerError)
	}

}
