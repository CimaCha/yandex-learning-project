package main

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"net/http"
	"strings"
)

type zlibWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func defaultHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<html><body>"+strings.Repeat("Hello, world<br>", 20)+"</body></html>")
}

func (w zlibWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func deflateHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает gzip-сжатие
		// это упрощённый пример. В реальном приложении следует проверять все
		// значения r.Header.Values("Accept-Encoding") и разбирать строку
		// на составные части, чтобы избежать неожиданных результатов
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			next.ServeHTTP(w, r)
			return
		}

		// создаём gzip.Writer поверх текущего w
		zl, err := zlib.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer zl.Close()

		w.Header().Set("Content-Encoding", "deflate")
		// передаём обработчику страницы переменную типа zlibWriter для вывода данных
		next.ServeHTTP(zlibWriter{ResponseWriter: w, Writer: zl}, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandle)
	err := http.ListenAndServe(":3000", deflateHandle(mux))
	if err != nil {
		panic(err)
	}
}
