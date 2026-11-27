package main

import (
	"fmt"
	"regexp"
)

func main() {
	emails := []string{"$€§@yandex.com", "ivan@mail.ru", "john@gmailyahoo", "fedor@gmail.com", "stepan@yahoo.com", "commanderpike@gmail.com", "greta@abcd@gmail_yahoo.com", "abc.def@mail.c", "abc.def@mail#archive.com", "abc.def@mail.org"}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	for _, email := range emails {
		fmt.Printf("%v is valid email: %t\n", email, re.MatchString(email))
	}
}
