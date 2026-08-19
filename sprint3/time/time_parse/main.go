package main

import (
	"fmt"
	"time"
)

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	fmt.Println(time.Parse(time.RFC3339, currentTimeStr))
}
