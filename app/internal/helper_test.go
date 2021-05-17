package internal

import (
	"testing"
	"strconv"
)

func TestParseFloat(t *testing.T) {
    result := ParseFloat("5.0")
    if result != 5.0 {
       t.Errorf("ParseFloat was incorrect, got:," + strconv.FormatFloat(result, 'E', -1, 64) + " want: 5")
    }
}

