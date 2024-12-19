package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		if root1 < root2 {
			fmt.Printf("%.3f %.3f\n", root1, root2)
		} else {
			fmt.Printf("%.3f %.3f\n", root2, root1)
		}
	} else if discriminant == 0 {
		root := -b / (2 * a)
		fmt.Printf("%.3f %.3f\n", root, root)
	} else {
		fmt.Println("0 0")
	}
}

func main() {
	SqRoots()
}
