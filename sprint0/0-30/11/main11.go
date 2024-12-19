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
		fmt.Println("Некорректный ввод")
	}

	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Некорректный ввод")
	}

	var final_number int

	for i := 1; i <= number; i++ {
		if i%2 == 1 {
			final_number += i
		}
	}
	fmt.Println(final_number)
}
