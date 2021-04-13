package service

import (
	"app/internal"
	"strconv"
	"time"
)

// DateEqual Check if two dates are equal
func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// SumUpLoadAmountsofTransactionRecords for to add all load amounts for transaction records
func SumUpLoadAmountsofTransactionRecords(loadTransactionRecords []internal.LoadTransactionRecord) float64 {
	result := 0.00
	for _, record := range loadTransactionRecords {
		loadAmount, _ := strconv.ParseFloat(record.LoadAmount, 64)
		result = result + loadAmount
	}
	return result
}
