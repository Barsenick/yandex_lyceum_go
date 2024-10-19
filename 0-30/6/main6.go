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
		fmt.Println("Некорректный ввод.")
	}

	if number < 0 {
		fmt.Println("Неккоректный ввод.")
	} else if number >= 1000 {
		fmt.Println("Число больше или равно 1000")
	} else if number < 10 {
		fmt.Println("Число меньше 10")
	} else if number < 100 {
		fmt.Println("Число меньше 100")
	} else if number < 1000 {
		fmt.Println("Число меньше 1000")
	}
}
