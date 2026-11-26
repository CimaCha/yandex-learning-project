package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	src := `/* Тестовый пакет */
package main

// Double умножает значение на 2.
func Double(i int) int {
    return i*2
}

func main() {
   // умножаем в цикле
   for i := 1; i < 5; i++ {
      fmt.Println(Double(i))
   }
}`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, comment := range f.Comments {
		for _, c := range comment.List {
			fmt.Println(fset.Position(c.Slash), c.Text)
		}
	}
}
