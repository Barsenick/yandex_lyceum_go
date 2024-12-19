package main

import (
	"fmt"
	"net/http"
)

var fibNumbers = []int{0, 1}
var requests = 0

func FibbonaciHandler(w http.ResponseWriter, r *http.Request) {
	if requests == 0 {
		fmt.Fprintf(w, "%d", fibNumbers[0])
	} else if requests == 1 {
		fmt.Fprintf(w, "%d", fibNumbers[1])
	} else {
		nextFib := fibNumbers[len(fibNumbers)-1] + fibNumbers[len(fibNumbers)-2]
		fibNumbers = append(fibNumbers, nextFib)
		fmt.Fprintf(w, "%d", fibNumbers[len(fibNumbers)-1])
	}
	requests++
}

func main() {
	http.HandleFunc("/", FibbonaciHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
