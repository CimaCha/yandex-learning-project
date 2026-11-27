package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Order struct {
	ID        int64
	CreatedAt int64
}

func main() {
	orders := []Order{
		{ID: 2, CreatedAt: 1735678801},
		{ID: 15, CreatedAt: 1735678805},
		{ID: 8, CreatedAt: 1735678802},
		{ID: 9, CreatedAt: 1735678803},
		{ID: 1, CreatedAt: 1735678800},
		{ID: 11, CreatedAt: 1735678805},
		{ID: 5, CreatedAt: 1735678802},
		{ID: 20, CreatedAt: 1735678808},
		{ID: 21, CreatedAt: 1735678810},
		{ID: 17, CreatedAt: 1735678812},
		{ID: 24, CreatedAt: 1735678811},
	}

	// Отсортируйте orders с использованием подходящего пакета из стандартной библиотеки.

	slices.SortFunc(orders, func(a Order, b Order) int {
		if a.CreatedAt == b.CreatedAt {
			return cmp.Compare(a.ID, b.ID)
		}
		return cmp.Compare(a.CreatedAt, b.CreatedAt)
	})

	fmt.Print(orders)
}
