package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: Gopher
values:
- 11
- 22
- 33
`

func main() {
	// вставьте недостающий код
	// 1) десериализуйте yamlData в переменную типа Data
	// 2) преобразуйте полученную переменную в TOML
	// 3) выведите в консоль результат
	// ...
	var data Data
	err := yaml.Unmarshal([]byte(yamlData), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	result, err := toml.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
