package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}
	response, err := client.Get("http://ya.ru")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.Copy(io.Discard, response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(strconv.FormatInt(body, 10))
	err = response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
