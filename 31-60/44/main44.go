package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

type Log struct {
	LogLevel LogLevel
}

func (l Log) Log(message string) {
	switch l.LogLevel {
	case Error:
		fmt.Printf("ERROR: %s\n", message)
	case Info:
		fmt.Printf("INFO: %s\n", message)
	default:
		fmt.Println("Invalid log level")
	}

}
