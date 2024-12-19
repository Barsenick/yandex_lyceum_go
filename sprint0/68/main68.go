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
	expression := "5+5"
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
		'(': 3,
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

		if r == ')' {
			ii := 0
			for r != '(' && 1-ii > 0 {
				nums = append(nums, operationsStack[len(operationsStack)-1-ii])
				ii++
			}
			continue
		}

		if r == '+' || r == '-' || r == '*' || r == '/' || r == '(' {
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

	left := make([]interface{}, 0)
	for _, r := range nums {
		if r != '(' && r != ')' {
			left = append(left, r)
		}
	}
	right := make([]float64, 0)
	operations := make([]rune, 0)

	for _, rr := range left {
		res, ok := rr.(float64)

		if ok {
			if len(operations) == 0 {
				right = append(right, res)
			}

			continue
		} else {
			operations = append(operations, rr.(rune))
		}
		evaluate(&right, &operations[len(operations)-1])
	}

	return result, nil
}

func evaluate(right *[]float64, operation *rune) (float64, error) {
	if len(*right) < 2 {
		return 0, ErrInvalidInput
	}

	res := (*right)[len(*right)-2]

	switch *operation {
	case '+':
		res += (*right)[len(*right)-1]
	case '-':
		res -= (*right)[len(*right)-1]
	case '*':
		res *= (*right)[len(*right)-1]
	case '/':
		if (*right)[len(*right)-1] == 0 {
			return 0, ErrDivisionByZero
		}
		res /= (*right)[len(*right)-1]
	default:
		return 0, ErrInvalidInput
	}

	*right = (*right)[:len(*right)-2]
	return res, nil
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
