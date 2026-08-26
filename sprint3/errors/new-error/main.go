package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// LabelError описывает ошибку с дополнительной меткой.
type LabelError struct {
	Label string // метка должна быть в верхнем регистре
	Err   error
}

// добавьте методы Error() и NewLabelError(label string, err error)

func (le *LabelError) Error() string {
	return fmt.Sprintf("[%v] %v", le.Label, le.Err)
}

func NewLabelError(label string, err error) error {
	return &LabelError{
		Label: strings.ToUpper(label),
		Err:   err,
	}
}

// ...

// программа должна выводить правильное значение

// Unwrap() возвращает исходную ошибку.
func (le *LabelError) Unwrap() error {
	return le.Err
}

// ...

func main() {
	_, err := os.ReadFile("mytest.txt")
	if err != nil {
		err = NewLabelError("file", err)
	}
	fmt.Println(errors.Is(err, os.ErrNotExist), err)
	// должна выводить текст:
	// true [FILE] open mytest.txt: no such file or directory
}
