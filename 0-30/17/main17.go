package main

import "fmt"

func IsPowerOfTwoRecursive(N int) {
	if N == 1 {
		fmt.Println("YES")
		return
	} else if N < 1 {
		fmt.Println("NO")
	} else {
		IsPowerOfTwoRecursive(N / 2)
	}
}

func main() {
	IsPowerOfTwoRecursive(16)
}
