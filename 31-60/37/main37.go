package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Adress string
}

func (p Person) Print() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:", p.Adress)
}
