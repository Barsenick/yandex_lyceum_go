package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

var (
	// Errors
	ErrDivisionByZero = errors.New("division by zero is not allowed")
	ErrInvalidInput   = errors.New("invalid input")
)

func main() {
	expression := "35 + (10 - 2 * 5) + (6 / 2 * 5 - 10 + 2) * (2 * 3)" //"2*(1+6)"
	result, err := Calc(expression)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Result: %.2f\n", result)
}

func Calc(expression string) (float64, error) {
	var result float64 = 0

	out_arr := make([]interface{}, 0)
	operationsStack := make([]rune, 0)

	Priority := map[rune]int{
		'*': 2,
		'/': 2,
		'+': 1,
		'-': 1,
	}

	i, j := 0, 0

	// удаляет все пробелы из строки

	newExpression := ""
	for _, r := range expression {
		if r != ' ' {
			newExpression += string(r)
		}
	}

	for _, r := range newExpression {
		// накапливаем цифры для преобразования в число
		if unicode.IsDigit(r) {
			j++
			// если последняя цифра в выражении - её нужно ниже добавить в out_arr
			if j < len(newExpression) {
				continue
			}
		}
		if i != j {
			num, err := strconv.ParseFloat(newExpression[i:j], 64)
			if err != nil {
				return result, err
			}
			out_arr = append(out_arr, num)
		}

		// накапливаем операции
		stack_len := len(operationsStack)
		if r == '+' || r == '-' || r == '*' || r == '/' {
			if stack_len == 0 {
				operationsStack = append(operationsStack, r)
			} else if Priority[r] <= Priority[operationsStack[stack_len-1]] {
				out_arr = append(out_arr, operationsStack[stack_len-1])
				operationsStack[stack_len-1] = r
			} else {
				operationsStack = append(operationsStack, r)
			}
		} else if r == '(' {
			operationsStack = append(operationsStack, r)
		} else if r == ')' {
			k := stack_len
			for operationsStack[k-1] != '(' {
				k--
				out_arr = append(out_arr, operationsStack[k])
				operationsStack = slices.Delete(operationsStack, k, k+1)
			}
			operationsStack = slices.Delete(operationsStack, k-1, k)
		}

		// пропускаем операторы и скобки для преобразования следующих чисел
		j++
		i = j
	}

	for i := range operationsStack {
		out_arr = append(out_arr, operationsStack[len(operationsStack)-1-i])
	}
	fmt.Println(out_arr)

	// ниже считает неправильно
	left := out_arr
	right := make([]interface{}, 0)

	for _, rr := range left {
		res, ok := rr.(float64)

		if ok {
			if len(right) > 1 {
				lastElem := right[len(right)-1]
				if _, ok := lastElem.(float64); ok {
					right = append(right, res)
				} else {
					if lastElem == '+' {
						result = right[len(right)-2].(float64) + res
						right = slices.Delete(right, len(right)-2, len(right))
						right = append(right, result)

					} else if lastElem == '-' {
						result = right[len(right)-2].(float64) - res
						right = slices.Delete(right, len(right)-2, len(right))
						right = append(right, result)

					} else if lastElem == '*' {
						result = right[len(right)-2].(float64) * res
						right = slices.Delete(right, len(right)-2, len(right))
						right = append(right, result)

					} else if lastElem == '/' {
						if res == 0 {
							return 0, ErrDivisionByZero
						}
						result = right[len(right)-2].(float64) / res
						right = slices.Delete(right, len(right)-2, len(right))
						right = append(right, result)
					}
				}
			} else {
				right = append(right, res)
			}
			continue
		}
		if len(right) >= 2 {
			if rr == '+' {
				result = right[len(right)-2].(float64) + right[len(right)-1].(float64)
				right = slices.Delete(right, len(right)-2, len(right))
				right = append(right, result)

			} else if rr == '-' {
				result = right[len(right)-2].(float64) - right[len(right)-1].(float64)
				right = slices.Delete(right, len(right)-2, len(right))

				right = append(right, result)
			} else if rr == '*' {
				result = right[len(right)-2].(float64) * right[len(right)-1].(float64)
				right = slices.Delete(right, len(right)-2, len(right))
				right = append(right, result)

			} else if rr == '/' {
				if right[len(right)-1] == 0 {
					return 0, ErrDivisionByZero
				}
				result = right[len(right)-2].(float64) / right[len(right)-1].(float64)
				right = slices.Delete(right, len(right)-2, len(right))
				right = append(right, result)
			}
		} else {
			right = append(right, rr)
		}
	}

	return result, nil
}

/*
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
*/
