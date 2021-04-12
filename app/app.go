package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	service "app/internal/service"

	_ "github.com/lib/pq"
)

func main() {

	dbConnectionString := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"
	db, dbOpenErr := sql.Open("postgres", dbConnectionString)
	if dbOpenErr != nil {
		log.Fatal("Failed to open a DB connection: ", dbOpenErr)
	}
	defer db.Close()

	service := service.New(db)

	file, fileOpenErr := os.Open(os.Getenv("INPUT_FILE"))
	if fileOpenErr != nil {
		log.Fatal(fileOpenErr)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		message := scanner.Text()
		var response string
		duplicateRecord, responseObj, loadFundsErr := service.LoadFunds(message)
		if loadFundsErr != nil {
			log.Fatal(loadFundsErr)
		}
		if duplicateRecord {
			continue
		}
		_ = json.Unmarshal(responseObj, &response)
		fmt.Println(response)
	}
	if scannerError := scanner.Err(); scannerError != nil {
		log.Fatal(scannerError)
	}

}
