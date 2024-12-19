package main

import (
	"fmt"
	"os"
	"strconv"
)

func run() error {
	// Проверяем, что передано достаточное количество аргументов
	if len(os.Args) != 4 {
		return fmt.Errorf("usage: %s <width> <height> <fill percentage>", os.Args[0])
	}

	// Преобразуем аргументы в числа
	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return fmt.Errorf("invalid width: %s", os.Args[1])
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return fmt.Errorf("invalid height: %s", os.Args[2])
	}

	fillPercentage, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return fmt.Errorf("invalid fill percentage: %s", os.Args[3])
	}

	// Формируем строку для записи в файл
	configStr := fmt.Sprintf("%dx%d %d%%", width, height, fillPercentage)

	// Записываем конфигурацию в файл
	err = os.WriteFile("config.txt", []byte(configStr), 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}
