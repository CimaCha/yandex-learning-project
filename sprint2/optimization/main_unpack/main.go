package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// LengthHandle возвращает размер распакованных данных.
func LengthHandle(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Length: %d", len(data))
		return
	}
	// создаём *gzip.Reader, который будет читать тело запроса
	// и распаковывать его
	gz, err := gzip.NewReader(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// закрытие gzip-читателя опционально, так как все данные уже прочитаны и
	// текущая реализация не требует закрытия, тем не менее лучше это делать -
	// некоторые реализации могут рассчитывать на закрытие читателя
	// gz.Close() не вызывает закрытия r.Body - это будет сделано позже, http-сервером
	defer gz.Close()

	// при чтении вернётся распакованный слайс байт
	body, err := io.ReadAll(gz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Length: %d", len(body))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", LengthHandle)
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
