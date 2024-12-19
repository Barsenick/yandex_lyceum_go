package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var (
	// Errors
	ErrDivisionByZero = errors.New("division by zero is not allowed")
	ErrInvalidInput   = errors.New("invalid input")
)

func main() {
	expression := "1+6"
	result, err := Calc(expression)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Result: %.2f\n", result)
}

func Calc(expression string) (float64, error) {
	var result float64 = 0

	nums := make([]interface{}, 0)
	operationsStack := make([]rune, 0)

	Priority := map[rune]int{
		'*': 2,
		'/': 2,
		'+': 1,
		'-': 1,
	}

	i, j := 0, 0

	// удалить все пробелы из строки

	for _, r := range expression {
		if unicode.IsDigit(r) {
			j++
			if j == len(expression) {
				strNum, _ := getFirstNumber(expression[i:j])
				num, err := strconv.ParseFloat(strNum, 64)
				i = j

				if err != nil {
					return result, err
				}
				nums = append(nums, num)
			}
			continue
		}

		strNum, _ := getFirstNumber(expression[i:j])
		num, err := strconv.ParseFloat(strNum, 64)
		i = j

		if err != nil {
			return result, err
		}
		nums = append(nums, num)

		if r == '+' || r == '-' || r == '*' || r == '/' {
			if len(operationsStack) == 0 {
				operationsStack = append(operationsStack, r)
			} else if Priority[r] <= Priority[operationsStack[len(operationsStack)-1]] {
				nums = append(nums, operationsStack[len(operationsStack)-1])
				operationsStack[len(operationsStack)-1] = r
			} else {
				operationsStack = append(operationsStack, r)
			}
			i++
			j = i
			continue
		}
	}

	for i := range operationsStack {
		nums = append(nums, operationsStack[len(operationsStack)-1-i])
	}

	left := nums
	right := make([]float64, 0)

	for _, rr := range left {
		res, ok := rr.(float64)

		if ok {
			right = append(right, res)
			continue
		}

		if rr == '+' {
			result = right[len(right)-2] + right[len(right)-1]
			right = []float64{right[0], result}
			continue
		}
		if rr == '-' {
			result = right[len(right)-2] - right[len(right)-1]
			right = []float64{right[0], result}
			continue
		}
		if rr == '*' {
			result = right[len(right)-2] * right[len(right)-1]
			right = []float64{right[0], result}
			continue
		}
		if rr == '/' {
			if right[len(right)-1] == 0 {
				return 0, ErrDivisionByZero
			}
			result = right[len(right)-2] / right[len(right)-1]
			right = []float64{right[0], result}
			continue
		}
	}

	return result, nil
}

func getFirstNumber(expr string) (string, int) {
	if len(expr) == 0 {
		return "", 0
	}
	if !unicode.IsDigit([]rune(expr)[0]) {
		return string([]rune(expr)[0]), 0
	}

	pos := 0
	for pos, val := range expr {
		if !unicode.IsDigit(val) {
			return expr[:pos], pos
		}
	}
	return expr, pos
}
