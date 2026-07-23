package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"net/http"
	"yandex-learning-project/sprint1/request/model"
)

func main() {
	url := "https://jsonplaceholder.typicode.com"

	client := gentleman.New()

	client.URL(url)

	request := client.Request()
	request.Path("/users")
	request.Method(http.MethodGet)

	result, err := request.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !result.Ok {
		fmt.Printf("Invalid server response: %d\n", result.StatusCode)
		return
	}

	users := &[]model.User{}
	err = result.JSON(users)
	if err != nil {
		fmt.Printf("JSON Parse error: %s\n", err)
		return
	}

	fmt.Print("Usernames: ")
	for _, user := range *users {
		fmt.Printf("%+v ", user.Username)
	}
	//result.String()
	// если выбрали resty, используйте SetResult(&users)
	// для получения результата сразу в виде массива
	// ...
}
