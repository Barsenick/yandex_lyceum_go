package main

import (
	"fmt"
	"net/http"
)

var fibNumbers = []int{0, 1}
var requests = 0

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
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

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requests)
}

func main() {
	http.HandleFunc("/", FibonacciHandler)
	http.HandleFunc("/metrics", MetricsHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
