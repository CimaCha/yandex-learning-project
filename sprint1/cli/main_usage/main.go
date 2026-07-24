package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

func main() {
	flag.Usage = func() {
		_, err := fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [options]\nversion: %s\n", os.Args[0], version)
		if err != nil {
			return
		}
		flag.PrintDefaults()
	}
	flag.Parse()
}
