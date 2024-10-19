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

	fmt.Println("Введите два числа: ")

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

	if num1 > num2 {
		fmt.Println("Первое число больше второго")

	} else if num1 < num2 {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}
