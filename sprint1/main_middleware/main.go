package main

import "net/http"

// middleware принимает параметром Handler и возвращает тоже Handler.
func middleware(next http.Handler) http.Handler {
	// получаем Handler приведением типа http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// здесь пишем логику обработки
		// например, разрешаем запросы cross-domain
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// ...
		// замыкание: используем ServeHTTP следующего хендлера
		next.ServeHTTP(w, r)
	})
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет"))
}
func main() {
	http.Handle("/", middleware(http.HandlerFunc(rootHandle)))
	//...
}
