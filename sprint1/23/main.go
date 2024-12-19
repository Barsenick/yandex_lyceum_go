package main

import (
	"fmt"
	"net/http"
)

func isLatin(input string) bool {
	for _, r := range input {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		name := query.Get("name")
		if !isLatin(name) {
			name = "dirty hacker"
		}
		query.Set("name", name)
		r.URL.RawQuery = query.Encode()
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		name := query.Get("name")
		if name == "" {
			name = "stranger"
		}
		query.Set("name", name)
		r.URL.RawQuery = query.Encode()
		next.ServeHTTP(w, r)
	})
}

func StrangerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", Sanitize(SetDefaultName(StrangerHandler)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
