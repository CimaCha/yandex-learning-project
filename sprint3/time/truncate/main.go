package main

import (
	"fmt"
	"time"
)

func main() {
	var today time.Time
	today = time.Now().Truncate(24 * time.Hour)
	fmt.Println(today)
}
