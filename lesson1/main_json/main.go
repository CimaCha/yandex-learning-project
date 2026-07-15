package main_json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"
	for k, v := range req.URL.Query() {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Form parameters ===============\r\n"
	err := req.ParseForm()
	if err != nil {
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	_, err = res.Write([]byte(body))
	if err != nil {
		return
	}
}

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// собираем данные
	subj := Subj{"Milk", 50}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusOK)
	// пишем тело ответа
	w.Write(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/main`, mainPage)
	mux.HandleFunc(`/json`, JSONHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
