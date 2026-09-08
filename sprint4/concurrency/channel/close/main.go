package main

import (
	"fmt"
	"sync"
)

func main() {
	chIn := make(chan int)
	chOut := make(chan int)
	quit := make(chan struct{})
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 15; i++ {
			chIn <- i
		}
		close(quit)
	}()
	go func() {
		var x int
		for {
			select {
			case x = <-chIn:
				chOut <- x * 2
			case <-quit:
				wg.Done()
				return
			}
		}
	}()
	wg.Add(1)
	go func() {
		for {
			fmt.Printf("%d ", <-chOut)
		}
	}()
	<-quit
	wg.Wait()
}
