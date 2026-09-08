package main

import (
	"bytes"
	"fmt"
	"sync"
)

type StringMerger struct {
	buffPool sync.Pool
}

func NewStringMerger() *StringMerger {
	return &StringMerger{
		buffPool: sync.Pool{
			// Функция New сработает, если в пуле нет объекта
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (s *StringMerger) Merge(strings ...string) string {
	buff := s.buffPool.Get().(*bytes.Buffer)
	// В конце отправляем буфер обратно в пул, чтобы переиспользовать
	defer s.buffPool.Put(buff)

	// Сбрасываем буфер перед использованием, но выделенная память сохранится
	buff.Reset()

	for _, s := range strings {
		buff.WriteString(s)
	}
	return buff.String()
}

func main() {
	stringConcat := NewStringMerger()

	firstString := stringConcat.Merge("Foo", "Bar", "Buzz")
	secondString := stringConcat.Merge("This is ", "a test line.")
	thirdString := stringConcat.Merge("Hello", ", ", "World!")

	_, _ = firstString, secondString
	fmt.Println(thirdString)
}
