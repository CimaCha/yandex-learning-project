package main

import "net/http"

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`Главная страница`))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`Это страница Api`))
}

func main() {
	http.HandleFunc(`/`, mainPage)
	http.HandleFunc(`/api`, apiPage)

	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}
