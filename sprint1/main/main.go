package main

import (
	"fmt"
	"io"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`Главная страница`))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`Это страница Api`))
}

func WriteHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(`/`, mainPage)
	mux.HandleFunc(`/api/`, apiPage)
	mux.HandleFunc(`/write/`, WriteHandle)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
