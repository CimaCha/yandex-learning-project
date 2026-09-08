package main

import "fmt"

func main() {
	c := gen(2, 3)
	out := square(c)

	for res := range out {
		fmt.Println(res)
	}
}

// реализация генератора gen здесь
func gen(nums ...int) chan int {
	inputCh := make(chan int)

	// горутина, в которой отправляем в канал  inputCh данные
	go func() {
		// как отправители закрываем канал, когда всё отправим
		defer close(inputCh)

		// перебираем все данные в слайсе
		for _, num := range nums {
			inputCh <- num
		}
	}()

	// возвращаем канал для данных
	return inputCh
}

// реализация square здесь
func square(in chan int) chan int {
	squareRes := make(chan int)

	go func() {
		defer close(squareRes)

		for num := range in {
			squareRes <- num * num
		}
	}()

	return squareRes
}
