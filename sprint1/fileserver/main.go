package main

import (
	"net/http"
)

func openfileHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.ServeFile(w, req, "./sprint1/fileserver/main.go")
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./"))
	mux.Handle(`/golang/`, http.StripPrefix(`/golang/`, fs))
	mux.HandleFunc("/", openfileHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
