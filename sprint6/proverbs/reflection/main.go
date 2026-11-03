package main

import (
	"io"
	"os"
)

func ConditionalWrite(w io.Writer) (int, error) {
	if f, ok := w.(*os.File); ok {

	}
}
