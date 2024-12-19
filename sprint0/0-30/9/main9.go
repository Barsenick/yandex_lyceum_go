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

	var final_number int

	for i := 1; i <= number; i++ {
		final_number += i
	}
	fmt.Println(final_number)
}
