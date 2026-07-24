package main

import (
	"fmt"
	"os"
)

func main() {
	// первый аргумент — имя запущенного файла
	fmt.Printf("Command: %v\n", os.Args[0])
	// выведем остальные параметры
	for i, v := range os.Args[1:] {
		fmt.Println(i+1, v)
	}
}
