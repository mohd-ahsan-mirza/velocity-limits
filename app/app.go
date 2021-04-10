package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	dbConnectionString := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME")
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	file, fileOpenErr := os.Open(os.Getenv("INPUT_FILE"))
	if fileOpenErr != nil {
		log.Fatal(fileOpenErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scannerError := scanner.Err(); scannerError != nil {
		log.Fatal(scannerError)
	}

}
