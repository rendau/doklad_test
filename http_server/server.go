package main

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, World!"))
}

func main() {
	mux := getMux()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func getMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", HelloHandler)
	return mux
}
