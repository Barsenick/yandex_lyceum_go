package main

import (
	"fmt"

	"github.com/Barsenick/calculator/calc"
)

func main() {
	fmt.Print("Enter an expression: ")
	expression := ""
	fmt.Scan(&expression)
	result, err := calc.Calc(expression)
	fmt.Println()
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("Result: %v", result)
	}
}
