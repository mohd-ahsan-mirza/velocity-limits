package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

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
