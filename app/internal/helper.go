package internal

import (
	"log"
	"strconv"
	"time"
)

// ParseFloat parsing string to float
func ParseFloat(input string) float64 {
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// ParseInt parsing string to int
func ParseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// DateEqual Check if two dates are equal
func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// SumUpLoadAmountsofTransactionRecords for to add all load amounts for transaction records
func SumUpLoadAmountsofTransactionRecords(loadTransactionRecords []LoadTransactionRecord) float64 {
	result := 0.00
	for _, record := range loadTransactionRecords {
		loadAmount, _ := strconv.ParseFloat(record.LoadAmount, 64)
		result = result + loadAmount
	}
	return result
}
