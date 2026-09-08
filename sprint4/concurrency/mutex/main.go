package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		go func(v int) {
			mu.Lock()
			defer mu.Unlock()
			m[v] = 1
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(m))
}
