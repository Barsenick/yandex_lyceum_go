package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}
