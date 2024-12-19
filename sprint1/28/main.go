package main

import (
	"fmt"
	"net/http"
)

func answerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The answer is 42")
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	http.HandleFunc("/answer/", Authorization(answerHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
