package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("USER:", os.Getenv("DB_USER"))
}
