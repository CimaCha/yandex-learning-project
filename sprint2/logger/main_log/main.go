package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	// допишите код
	mylog := log.New(&buf, `my log `, log.LstdFlags|log.Lshortfile)
	mylog.Println("Hello, world!")
	mylog.Println("Queen of Hearts")

	fmt.Print(&buf)
	// должна вывести
	// mylog: Hello, world!
	// mylog: Goodbye
}
