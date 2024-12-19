package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input = scanner.Text()

	numbers := strings.Split(input, " ")

	if len(numbers) != 2 {
		fmt.Println("Чисел должно быть 2!")
	}

	num1, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Println(err)
	}
	num2, err := strconv.Atoi(numbers[1])
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println("Неверный формат ввода.")
	}

	if num1 > 0 && num2 > 0 {
		fmt.Println("Оба числа положительные")
	} else if num1 < 0 && num2 < 0 {
		fmt.Println("Оба числа отрицательные")
	} else if num1 == 0 || num2 == 0 {
		fmt.Println("Одно из чисел равно нулю")
	} else {
		fmt.Println("Одно число положительное, а другое отрицательное")
	}
}
