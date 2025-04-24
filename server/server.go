package main

import (
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(100)
	// 5 % failure ratee
	if number > 5 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
	if number >2 && number <= 5 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}
	if number <= 2 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
}
