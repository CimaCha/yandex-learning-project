package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := generator("Hello")
	for msg := range ch {
		fmt.Println(msg)
	}
}

// Тут ваш генератор
func generator(msg string) chan string {
	channel := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			channel <- fmt.Sprintf("%s %s", msg, strconv.Itoa(i))
		}
		close(channel)
	}()
	return channel
}
