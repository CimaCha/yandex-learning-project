package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ticker := time.NewTicker(2 * time.Second)
	for range 10 {
		t := <-ticker.C
		fmt.Println(int(t.Sub(start).Seconds()))
	}
	ticker.Stop()
}
