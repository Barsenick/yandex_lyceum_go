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

	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input = scanner.Text()

	numbers := strings.Split(input, " ")

	if len(numbers) != 3 {
		fmt.Println("Некорректный ввод")
		return
	}

	num1, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}
	num2, err := strconv.Atoi(numbers[1])
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}
	num3, err := strconv.Atoi(numbers[2])
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}

	if num1 == num2 && num2 == num3 {
		fmt.Println("Все числа равны")
	} else if num1 == num2 || num2 == num3 || num1 == num3 {
		fmt.Println("Два числа равны")
	} else {
		fmt.Println("Все числа разные")
	}
}
