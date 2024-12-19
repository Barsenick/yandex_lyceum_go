package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	json, _ := json.Marshal(Response{name})
	log.Println(string(json))
	fmt.Fprint(w, string(json))
}

func main() {
	http.HandleFunc("/", HelloHandler)

	http.ListenAndServe(":8080", nil)
}
