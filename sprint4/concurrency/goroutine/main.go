package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Ядер:", runtime.NumCPU())
	fmt.Println("Логических процессоров:", runtime.GOMAXPROCS(2),
		"Горутин:", runtime.NumGoroutine())
	go func() {
		time.Sleep(100 * time.Millisecond)
	}()
	fmt.Println("Логических процессоров:", runtime.GOMAXPROCS(0),
		"Горутин:", runtime.NumGoroutine())
}
