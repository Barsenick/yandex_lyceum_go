package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	var input string = scanner.Text()

	if input == "" {
		fmt.Println("Пожалуйста, введите число.")
	}

	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Некорректный ввод.")
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(input, "*", i, "=", number*i)
	}
}
