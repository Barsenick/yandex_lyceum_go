package main

import (
	"fmt"
	"net/http"
)

var fibNumbers = []int{}
var requestCount = 0

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	if len(fibNumbers) == 0 {
		fibNumbers = append(fibNumbers, 0)
		fmt.Fprintf(w, "%d", fibNumbers[len(fibNumbers)-1])
	} else if len(fibNumbers) == 1 {
		fibNumbers = append(fibNumbers, 1)
		fmt.Fprintf(w, "%d", fibNumbers[len(fibNumbers)-1])
	} else {
		nextFib := fibNumbers[len(fibNumbers)-1] + fibNumbers[len(fibNumbers)-2]
		fibNumbers = append(fibNumbers, nextFib)
		fmt.Fprintf(w, "%d", fibNumbers[len(fibNumbers)-1])
	}
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

func main() {
	http.HandleFunc("/", FibonacciHandler)
	http.HandleFunc("/metrics", MetricsHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
