package main

import (
	"fmt"
)

func main() {
	array := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := make([]byte, 5)
	// ограничили длину len слайса, но его вместимость cap не изменилась
	fmt.Println(cap(slice)) // 10
	_ = copy(slice, array[:5])
	// так работает
	fmt.Println(cap(slice)) // 5
	// теперь, поскольку вместительность слайса ограничена,
	// при вызове append ему будет выделен новый участок памяти
	slice = append(slice, 42)
	fmt.Println(slice) // [0 1 2 3 4 42]
	// защитим несущий массив array от неожидаемых изменений
	fmt.Println(array) // [0 1 2 3 4 5 6 7 8 9]
}
