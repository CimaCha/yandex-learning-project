package main

import (
	"fmt"
	"time"
)

func main() {
	// допишите код здесь
	birthday := time.Date(2093, 11, 26, 0, 0, 0, 0, time.Local)
	duration := time.Until(birthday)
	days := int(duration.Hours() / 24)
	fmt.Println(days)
}
