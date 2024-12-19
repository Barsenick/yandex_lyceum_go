package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input = scanner.Text()

	if input == "" {
		fmt.Println("Пожалуйста, введите число.")
	}

	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Неверный формат ввода.")
	}
	if number > 0 {
		if number%2 == 0 {
			fmt.Println("Число положительное и четное")
		} else {
			fmt.Println("Число положительное и нечетное")
		}
	} else {
		if number%2 == 0 {
			fmt.Println("Число отрицательное и четное")
		} else {
			fmt.Println("Число отрицательное и нечетное")
		}
	}
}
