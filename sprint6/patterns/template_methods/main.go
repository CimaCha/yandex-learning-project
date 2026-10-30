package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 48}, // John старший
		{"Michael", 17},
		{"John", 26}, // John младший
	}

	sort.Slice(people, func(i, j int) bool {
		if people[i].Name == people[j].Name {
			return people[i].Age < people[j].Age
		}
		return people[i].Name < people[j].Name
	})
	fmt.Println(people)
}
