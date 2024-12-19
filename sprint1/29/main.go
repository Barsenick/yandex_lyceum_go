package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var resultJSON string
var greet greeting
var greetingName string

type greeting struct {
	Greetings string `json:"greetings"`
	Name      string `json:"name"`
}

type result struct {
	Status string   `json:"status"`
	Result greeting `json:"result"`
}

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
			panic("unsupported name")
		}
		query.Set("name", name)
		greetingName = name
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
			greetingName = "stranger"
		}
		query.Set("name", name)
		r.URL.RawQuery = query.Encode()
		next.ServeHTTP(w, r)
	})
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprint(w, `{"status": "error", "result": {}}`)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	greet = greeting{"hello", greetingName}
	result := result{"ok", greet}
	jsonresult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	resultJSON = string(jsonresult)
	fmt.Fprint(w, resultJSON)
}

func main() {
	http.HandleFunc("/", HelloHandler)

	handler := RPC(SetDefaultName(Sanitize((HelloHandler))))
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
